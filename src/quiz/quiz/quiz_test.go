package quiz

import (
	"encoding/json"
	"strings"
	"testing"
)

const sample = `
title: Test Quiz
questions:
  - type: single
    text: Pick a
    options:
      - {id: a, text: A}
      - {id: b, text: B}
    answer: a
  - type: multiple
    text: Pick a and c
    points: 2
    options:
      - {id: a, text: A}
      - {id: b, text: B}
      - {id: c, text: C}
    answer: [a, c]
  - type: truefalse
    text: Go is compiled
    answer: "true"
  - type: text
    text: Keyword for a goroutine
    answer: ["go", "go statement"]
`

func parseSample(t *testing.T) *Quiz {
	t.Helper()
	q, err := Parse([]byte(sample))
	if err != nil {
		t.Fatalf("Parse() failed: %v", err)
	}
	return q
}

func TestParse(t *testing.T) {
	q := parseSample(t)
	if q.Title != "Test Quiz" {
		t.Errorf("Title = %q, want %q", q.Title, "Test Quiz")
	}
	if len(q.Questions) != 4 {
		t.Fatalf("got %d questions, want 4", len(q.Questions))
	}
	// A truefalse question gets its options filled in during validation.
	if got := len(q.Questions[2].Options); got != 2 {
		t.Errorf("truefalse options = %d, want 2", got)
	}
	// Points default to 1, but an explicit value is kept.
	if got := q.Questions[0].Points; got != 1 {
		t.Errorf("default points = %d, want 1", got)
	}
	if got := q.TotalPoints(); got != 5 {
		t.Errorf("TotalPoints() = %d, want 5", got)
	}
}

func TestParseRejectsBrokenQuizzes(t *testing.T) {
	tests := map[string]string{
		"no title":          "questions:\n  - {text: x, answer: a}",
		"no questions":      "title: T",
		"unknown type":      "title: T\nquestions:\n  - {type: essay, text: x, answer: a}",
		"missing answer":    "title: T\nquestions:\n  - {type: text, text: x}",
		"answer not option": "title: T\nquestions:\n  - type: single\n    text: x\n    options: [{id: a, text: A}, {id: b, text: B}]\n    answer: z",
		"too few options":   "title: T\nquestions:\n  - type: single\n    text: x\n    options: [{id: a, text: A}]\n    answer: a",
		"duplicate ids":     "title: T\nquestions:\n  - type: single\n    text: x\n    options: [{id: a, text: A}, {id: a, text: B}]\n    answer: a",
		"text with options": "title: T\nquestions:\n  - type: text\n    text: x\n    options: [{id: a, text: A}, {id: b, text: B}]\n    answer: a",
	}
	for name, src := range tests {
		t.Run(name, func(t *testing.T) {
			if _, err := Parse([]byte(src)); err == nil {
				t.Error("Parse() accepted an invalid quiz")
			}
		})
	}
}

func TestSanitizeHidesTheAnswer(t *testing.T) {
	q := parseSample(t)
	public := q.Questions[0].Sanitize()
	if public.Text != "Pick a" {
		t.Errorf("Text = %q", public.Text)
	}
	// PublicQuestion has no answer field at all; this test documents that the
	// fields a student may see survive sanitizing.
	if len(public.Options) != 2 {
		t.Errorf("options = %d, want 2", len(public.Options))
	}
}

func TestGrade(t *testing.T) {
	q := parseSample(t)
	tests := []struct {
		name     string
		question int
		given    []string
		want     bool
		points   int
	}{
		{"single correct", 0, []string{"a"}, true, 1},
		{"single wrong", 0, []string{"b"}, false, 0},
		{"single empty", 0, nil, false, 0},
		{"multiple exact", 1, []string{"a", "c"}, true, 2},
		{"multiple other order", 1, []string{"c", "a"}, true, 2},
		{"multiple incomplete", 1, []string{"a"}, false, 0},
		{"multiple superset", 1, []string{"a", "b", "c"}, false, 0},
		{"truefalse correct", 2, []string{"true"}, true, 1},
		{"truefalse wrong", 2, []string{"false"}, false, 0},
		{"text exact", 3, []string{"go"}, true, 1},
		{"text normalized", 3, []string{"  GO  "}, true, 1},
		{"text second alternative", 3, []string{"Go   Statement"}, true, 1},
		{"text wrong", 3, []string{"goroutine"}, false, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := q.Questions[tc.question].Grade(tc.given)
			if res.Correct != tc.want {
				t.Errorf("Correct = %v, want %v", res.Correct, tc.want)
			}
			if res.Points != tc.points {
				t.Errorf("Points = %d, want %d", res.Points, tc.points)
			}
		})
	}
}

func TestLoadDir(t *testing.T) {
	quizzes, problems := LoadDir("../quizzes")
	for _, p := range problems {
		t.Errorf("quiz in ../quizzes is broken: %v", p)
	}
	if len(quizzes) == 0 {
		t.Fatal("no quizzes loaded from ../quizzes")
	}
	for _, id := range SortedIDs(quizzes) {
		if quizzes[id].TotalPoints() <= 0 {
			t.Errorf("quiz %q has no points", id)
		}
	}
}

// TestSanitizeJSONFieldNames pins the wire format the TypeScript client reads.
// Without explicit json tags the nested structs serialize as "Data", "Lang"
// and "ID", which silently breaks the browser.
func TestSanitizeJSONFieldNames(t *testing.T) {
	q, err := Parse([]byte(`
title: T
questions:
  - type: single
    text: x
    code: {lang: go, source: "fmt.Println()"}
    image: {data: QUJD, mime: image/png, alt: a picture}
    options:
      - {id: a, text: A, image: {data: QUJD, mime: image/png}}
      - {id: b, text: B}
    answer: a
`))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(q.Questions[0].Sanitize())
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{`"type"`, `"text"`, `"points"`, `"lang"`, `"source"`, `"data"`, `"mime"`, `"alt"`, `"id"`, `"options"`} {
		if !strings.Contains(string(data), want) {
			t.Errorf("sanitized JSON is missing %s: %s", want, data)
		}
	}
	for _, unwanted := range []string{`"Data"`, `"MIME"`, `"Lang"`, `"ID"`, `"answer"`, `"Answer"`} {
		if strings.Contains(string(data), unwanted) {
			t.Errorf("sanitized JSON contains %s: %s", unwanted, data)
		}
	}
}
