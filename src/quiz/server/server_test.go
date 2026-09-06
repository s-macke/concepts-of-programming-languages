package server

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/session"
)

const sample = `
title: Test Quiz
questions:
  - type: single
    text: Pick a
    options: [{id: a, text: A}, {id: b, text: B}]
    answer: a
    explanation: because a
  - type: text
    text: Keyword
    answer: ["go"]
`

// newTestServer builds a server with a single quiz under the id "test".
func newTestServer(t *testing.T) http.Handler {
	t.Helper()
	q, err := quiz.Parse([]byte(sample))
	if err != nil {
		t.Fatalf("cannot parse test quiz: %v", err)
	}
	quizzes := map[string]*quiz.Quiz{"test": q}
	return New(quizzes, session.NewStore(time.Hour), os.DirFS(t.TempDir())).Routes()
}

// do performs a request and returns status and body.
func do(t *testing.T, h http.Handler, method, path string, body io.Reader, contentType string) (int, []byte) {
	t.Helper()
	req := httptest.NewRequest(method, path, body)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec.Code, rec.Body.Bytes()
}

// startSession creates a session and returns its id.
func startSession(t *testing.T, h http.Handler) string {
	t.Helper()
	status, body := do(t, h, "POST", "/api/sessions", strings.NewReader(`{"quizId":"test"}`), "application/json")
	if status != http.StatusCreated {
		t.Fatalf("POST /api/sessions = %d, body %s", status, body)
	}
	var info session.Info
	if err := json.Unmarshal(body, &info); err != nil {
		t.Fatalf("cannot decode session info: %v", err)
	}
	if info.NumQuestions != 2 || info.TotalPoints != 2 {
		t.Fatalf("info = %+v, want 2 questions and 2 points", info)
	}
	return info.SessionID
}

func TestListQuizzes(t *testing.T) {
	h := newTestServer(t)
	status, body := do(t, h, "GET", "/api/quizzes", nil, "")
	if status != http.StatusOK {
		t.Fatalf("status = %d", status)
	}
	var list []quizListEntry
	if err := json.Unmarshal(body, &list); err != nil {
		t.Fatal(err)
	}
	if len(list) != 1 || list[0].ID != "test" || list[0].NumQuestions != 2 {
		t.Errorf("list = %+v", list)
	}
}

// TestQuestionHidesTheAnswer is the central guarantee of the API: a student
// must not be able to read the answer out of the response in devtools.
func TestQuestionHidesTheAnswer(t *testing.T) {
	h := newTestServer(t)
	id := startSession(t, h)

	status, body := do(t, h, "GET", "/api/sessions/"+id+"/questions/1", nil, "")
	if status != http.StatusOK {
		t.Fatalf("status = %d, body %s", status, body)
	}
	var raw map[string]any
	if err := json.Unmarshal(body, &raw); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"answer", "Answer", "explanation", "Explanation"} {
		if _, ok := raw[forbidden]; ok {
			t.Errorf("question response contains %q: %s", forbidden, body)
		}
	}
	if bytes.Contains(body, []byte("because a")) {
		t.Errorf("question response leaks the explanation: %s", body)
	}
}

func TestAnswerGradesAndScores(t *testing.T) {
	h := newTestServer(t)
	id := startSession(t, h)

	status, body := do(t, h, "POST", "/api/sessions/"+id+"/answers/1", strings.NewReader(`{"answer":["a"]}`), "application/json")
	if status != http.StatusOK {
		t.Fatalf("status = %d, body %s", status, body)
	}
	var res quiz.Result
	if err := json.Unmarshal(body, &res); err != nil {
		t.Fatal(err)
	}
	if !res.Correct || res.Points != 1 || res.Explanation != "because a" {
		t.Errorf("result = %+v", res)
	}

	// A second attempt at the same question must not change the score.
	do(t, h, "POST", "/api/sessions/"+id+"/answers/1", strings.NewReader(`{"answer":["b"]}`), "application/json")

	status, body = do(t, h, "GET", "/api/sessions/"+id+"/result", nil, "")
	if status != http.StatusOK {
		t.Fatalf("status = %d", status)
	}
	var sum session.Summary
	if err := json.Unmarshal(body, &sum); err != nil {
		t.Fatal(err)
	}
	if sum.Points != 1 || sum.Correct != 1 || sum.Answered != 1 {
		t.Errorf("summary = %+v, want 1 point, 1 correct, 1 answered", sum)
	}
}

func TestNotFoundCases(t *testing.T) {
	h := newTestServer(t)
	id := startSession(t, h)

	tests := []struct {
		name, method, path, body string
		want                     int
	}{
		{"unknown quiz", "POST", "/api/sessions", `{"quizId":"nope"}`, http.StatusNotFound},
		{"unknown session", "GET", "/api/sessions/deadbeef/questions/1", "", http.StatusNotFound},
		{"question out of range", "GET", "/api/sessions/" + id + "/questions/99", "", http.StatusNotFound},
		{"question zero", "GET", "/api/sessions/" + id + "/questions/0", "", http.StatusBadRequest},
		{"question not a number", "GET", "/api/sessions/" + id + "/questions/x", "", http.StatusBadRequest},
		{"malformed answer body", "POST", "/api/sessions/" + id + "/answers/1", `{`, http.StatusBadRequest},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			status, _ := do(t, h, tc.method, tc.path, strings.NewReader(tc.body), "application/json")
			if status != tc.want {
				t.Errorf("status = %d, want %d", status, tc.want)
			}
		})
	}
}

func TestUpload(t *testing.T) {
	h := newTestServer(t)

	upload := func(t *testing.T, content string) (int, []byte) {
		t.Helper()
		var buf bytes.Buffer
		w := multipart.NewWriter(&buf)
		part, err := w.CreateFormFile("file", "quiz.yaml")
		if err != nil {
			t.Fatal(err)
		}
		part.Write([]byte(content))
		w.Close()
		return do(t, h, "POST", "/api/upload", &buf, w.FormDataContentType())
	}

	t.Run("valid file starts a session", func(t *testing.T) {
		status, body := upload(t, sample)
		if status != http.StatusCreated {
			t.Fatalf("status = %d, body %s", status, body)
		}
		var info session.Info
		if err := json.Unmarshal(body, &info); err != nil {
			t.Fatal(err)
		}
		if info.SessionID == "" || info.Title != "Test Quiz" {
			t.Errorf("info = %+v", info)
		}
		// The uploaded quiz is not published to everyone.
		_, list := do(t, h, "GET", "/api/quizzes", nil, "")
		if bytes.Count(list, []byte(`"id"`)) != 1 {
			t.Errorf("upload leaked into the quiz list: %s", list)
		}
	})

	t.Run("invalid file is rejected with a reason", func(t *testing.T) {
		status, body := upload(t, "title: Broken\nquestions: []")
		if status != http.StatusUnprocessableEntity {
			t.Fatalf("status = %d, body %s", status, body)
		}
		if !bytes.Contains(body, []byte("no questions")) {
			t.Errorf("error message is not helpful: %s", body)
		}
	})

	t.Run("missing file field", func(t *testing.T) {
		status, _ := do(t, h, "POST", "/api/upload", strings.NewReader(""), "multipart/form-data; boundary=x")
		if status != http.StatusBadRequest {
			t.Errorf("status = %d, want 400", status)
		}
	})
}
