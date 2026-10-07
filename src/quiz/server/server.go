// Package server exposes the quiz engine over HTTP.
//
// Solo questions reveal their answer only after submission. Group questions
// reveal answers and statistics only when the host closes the question.
package server

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/group"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/session"
)

// Server holds the quizzes read from disk and the running sessions.
type Server struct {
	quizzes   map[string]*quiz.Quiz
	sessions  *session.Store
	static    fs.FS
	groups    *group.Store
	publicURL string
}

// New creates a server serving the given quizzes and the web UI from static.
func New(quizzes map[string]*quiz.Quiz, sessions *session.Store, static fs.FS) *Server {
	return &Server{quizzes: quizzes, sessions: sessions, static: static, groups: group.NewStore(time.Hour)}
}

// Routes builds the HTTP handler. It uses the method and wildcard patterns
// introduced in Go 1.22, so no third party router is needed.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	s.groupRoutes(mux)

	mux.HandleFunc("GET /api/quizzes", s.handleListQuizzes)
	mux.HandleFunc("POST /api/sessions", s.handleCreateSession)
	mux.HandleFunc("POST /api/upload", s.handleUpload)
	mux.HandleFunc("GET /api/sessions/{id}", s.handleSessionInfo)
	mux.HandleFunc("GET /api/sessions/{id}/questions/{n}", s.handleQuestion)
	mux.HandleFunc("POST /api/sessions/{id}/answers/{n}", s.handleAnswer)
	mux.HandleFunc("GET /api/sessions/{id}/result", s.handleResult)

	// Everything else is the single page app.
	mux.Handle("/", http.FileServer(http.FS(s.static)))

	return logRequests(mux)
}

// quizListEntry describes one quiz on the start page.
type quizListEntry struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	NumQuestions int    `json:"numQuestions"`
}

func (s *Server) handleListQuizzes(w http.ResponseWriter, r *http.Request) {
	list := make([]quizListEntry, 0, len(s.quizzes))
	for _, id := range quiz.SortedIDs(s.quizzes) {
		q := s.quizzes[id]
		list = append(list, quizListEntry{
			ID:           id,
			Title:        q.Title,
			Description:  q.Description,
			NumQuestions: len(q.Questions),
		})
	}
	writeJSON(w, http.StatusOK, list)
}

func (s *Server) handleCreateSession(w http.ResponseWriter, r *http.Request) {
	var req struct {
		QuizID string `json:"quizId"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 4<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "malformed request body")
		return
	}
	q, ok := s.quizzes[req.QuizID]
	if !ok {
		writeError(w, http.StatusNotFound, "unknown quiz")
		return
	}
	s.startSession(w, q)
}

// startSession creates a session for q and answers with its Info.
func (s *Server) startSession(w http.ResponseWriter, q *quiz.Quiz) {
	sess, err := session.NewSolo(q)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "cannot start session")
		return
	}
	s.sessions.Add(sess)
	writeJSON(w, http.StatusCreated, sess.Info())
}

func (s *Server) handleSessionInfo(w http.ResponseWriter, r *http.Request) {
	sess, ok := s.lookup(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, sess.Info())
}

func (s *Server) handleQuestion(w http.ResponseWriter, r *http.Request) {
	sess, ok := s.lookup(w, r)
	if !ok {
		return
	}
	index, ok := questionIndex(w, r)
	if !ok {
		return
	}
	q, err := sess.Question(index)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, q)
}

func (s *Server) handleAnswer(w http.ResponseWriter, r *http.Request) {
	sess, ok := s.lookup(w, r)
	if !ok {
		return
	}
	index, ok := questionIndex(w, r)
	if !ok {
		return
	}
	var req struct {
		Answer []string `json:"answer"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 64<<10)).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "malformed request body")
		return
	}
	res, err := sess.Answer(index, req.Answer)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, res)
}

func (s *Server) handleResult(w http.ResponseWriter, r *http.Request) {
	sess, ok := s.lookup(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, sess.Summary())
}

// lookup resolves the {id} path value to a session, writing an error response
// if there is none.
func (s *Server) lookup(w http.ResponseWriter, r *http.Request) (session.Session, bool) {
	sess, ok := s.sessions.Get(r.PathValue("id"))
	if !ok {
		writeError(w, http.StatusNotFound, "unknown or expired session")
		return nil, false
	}
	return sess, true
}

// questionIndex parses the {n} path value, which is one based in the URL.
func questionIndex(w http.ResponseWriter, r *http.Request) (int, bool) {
	n, err := strconv.Atoi(r.PathValue("n"))
	if err != nil || n < 1 {
		writeError(w, http.StatusBadRequest, "question number must be a positive integer")
		return 0, false
	}
	return n - 1, true
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("cannot write response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// logRequests is a minimal middleware, mostly so that a lecture demo shows
// what the browser actually does. Log route templates only: real paths contain
// session credentials or room identifiers. Never log bodies, headers or IPs.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		log.Printf("%s %s", r.Method, r.Pattern)
	})
}
