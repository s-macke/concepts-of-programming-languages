package server

import (
	"bytes"
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func groupReq(t *testing.T, h http.Handler, method, path, token, body string, want int) map[string]any {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != want {
		t.Fatalf("%s %s: %d %s", method, path, rec.Code, rec.Body.String())
	}
	var data map[string]any
	if e := json.Unmarshal(rec.Body.Bytes(), &data); e != nil {
		t.Fatal(e)
	}
	return data
}
func TestGroupHTTP(t *testing.T) {
	h := newTestServer(t)
	var logs bytes.Buffer
	old := log.Writer()
	log.SetOutput(&logs)
	defer log.SetOutput(old)
	room := groupReq(t, h, "POST", "/api/groups", "", `{"quizId":"test"}`, 201)
	id := room["id"].(string)
	host := room["token"].(string)
	path := "/api/groups/" + id
	joined := groupReq(t, h, "POST", path+"/join", "", "{}", 201)
	p := joined["token"].(string)
	groupReq(t, h, "GET", path, "", "", 403)
	groupReq(t, h, "POST", path+"/actions", p, `{"action":"start","revision":0}`, 403)
	groupReq(t, h, "POST", path+"/actions", host, `{"action":"start","revision":0}`, 200)
	ack := groupReq(t, h, "POST", path+"/answers/1", p, `{"answer":["a"]}`, 200)
	if len(ack) != 1 || ack["accepted"] != true {
		t.Fatal(ack)
	}
	st := groupReq(t, h, "GET", path, p, "", 200)
	if st["correctAnswer"] != nil || st["personal"] != nil || st["stats"] != nil {
		t.Fatal(st)
	}
	groupReq(t, h, "POST", path+"/actions", host, `{"action":"reveal","revision":1}`, 200)
	st = groupReq(t, h, "GET", path, p, "", 200)
	if st["personal"] == nil || st["stats"] == nil {
		t.Fatal(st)
	}
	groupReq(t, h, "POST", path+"/actions", host, `{"action":"next","revision":1}`, 409)
	groupReq(t, h, "POST", path+"/actions", host, `{"action":"delete","revision":2}`, 200)
	groupReq(t, h, "GET", path, p, "", 404)
	for _, secret := range []string{id, host, p} {
		if strings.Contains(logs.String(), secret) {
			t.Fatalf("credential or room ID leaked into logs: %s", logs.String())
		}
	}
}
func TestGroupUpload(t *testing.T) {
	h := newTestServer(t)
	var b bytes.Buffer
	m := multipart.NewWriter(&b)
	f, e := m.CreateFormFile("file", "test.yaml")
	if e != nil {
		t.Fatal(e)
	}
	f.Write([]byte(sample))
	m.Close()
	code, data := do(t, h, "POST", "/api/groups/upload", &b, m.FormDataContentType())
	if code != 201 {
		t.Fatalf("%d %s", code, data)
	}
}
