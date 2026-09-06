package server

import (
	"io"
	"net/http"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

// handleUpload accepts a quiz YAML file as multipart form field "file" and
// immediately starts a session for it. The quiz is only ever kept in memory,
// inside that session, and disappears with it: uploading is a way to try a
// quiz out, not a way to publish one.
func (s *Server) handleUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, quiz.MaxFileSize)

	file, _, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, "expected a multipart form with a 'file' field")
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		writeError(w, http.StatusRequestEntityTooLarge, "quiz file is too large")
		return
	}
	q, err := quiz.Parse(data)
	if err != nil {
		// The parse error names the offending question, so it is worth showing.
		writeError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	s.startSession(w, q)
}
