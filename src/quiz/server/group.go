package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/group"
	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

// ConfigureGroups must be called before serving requests. Its janitor stops
// with ctx; all room state is discarded on process exit.
func (s *Server) ConfigureGroups(ctx context.Context, ttl time.Duration, publicURL string) {
	s.groups = group.NewStore(ttl)
	s.publicURL = strings.TrimRight(publicURL, "/")
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.groups.Expire()
			}
		}
	}()
}
func groupError(w http.ResponseWriter, err error) {
	code := http.StatusConflict
	if errors.Is(err, group.ErrMissing) {
		code = 404
	}
	if errors.Is(err, group.ErrAuth) {
		code = 403
	}
	if errors.Is(err, group.ErrAnswer) {
		code = 400
	}
	writeError(w, code, err.Error())
}
func groupBody(w http.ResponseWriter, r *http.Request, v any) bool {
	if json.NewDecoder(http.MaxBytesReader(w, r.Body, 64<<10)).Decode(v) != nil {
		writeError(w, 400, "malformed request body")
		return false
	}
	return true
}
func (s *Server) groupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/groups", s.createGroup)
	mux.HandleFunc("POST /api/groups/upload", s.uploadGroup)
	mux.HandleFunc("POST /api/groups/{id}/join", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		t, err := s.groups.Join(r.PathValue("id"))
		if err != nil {
			groupError(w, err)
			return
		}
		writeJSON(w, 201, map[string]string{"token": t})
	})
	mux.HandleFunc("GET /api/groups/{id}", func(w http.ResponseWriter, r *http.Request) {
		st, err := s.groups.State(r.PathValue("id"), bearer(r))
		if err != nil {
			groupError(w, err)
			return
		}
		w.Header().Set("Cache-Control", "no-store")
		writeJSON(w, 200, st)
	})
	mux.HandleFunc("POST /api/groups/{id}/answers/{n}", func(w http.ResponseWriter, r *http.Request) {
		i, ok := questionIndex(w, r)
		if !ok {
			return
		}
		var body struct {
			Answer []string `json:"answer"`
		}
		if !groupBody(w, r, &body) {
			return
		}
		if err := s.groups.Answer(r.PathValue("id"), bearer(r), i+1, body.Answer); err != nil {
			groupError(w, err)
			return
		}
		writeJSON(w, 200, map[string]bool{"accepted": true})
	})
	mux.HandleFunc("POST /api/groups/{id}/actions", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Action   string `json:"action"`
			Revision int    `json:"revision"`
		}
		if !groupBody(w, r, &body) {
			return
		}
		if err := s.groups.Action(r.PathValue("id"), bearer(r), body.Action, body.Revision); err != nil {
			groupError(w, err)
			return
		}
		writeJSON(w, 200, map[string]bool{"ok": true})
	})
}
func bearer(r *http.Request) string {
	value, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !ok {
		return ""
	}
	return value
}
func (s *Server) startGroup(w http.ResponseWriter, q *quiz.Quiz) {
	id, t := s.groups.Create(q)
	w.Header().Set("Cache-Control", "no-store")
	writeJSON(w, 201, map[string]string{"id": id, "token": t, "publicURL": s.publicURL})
}
func (s *Server) createGroup(w http.ResponseWriter, r *http.Request) {
	var body struct {
		QuizID string `json:"quizId"`
	}
	if !groupBody(w, r, &body) {
		return
	}
	q := s.quizzes[body.QuizID]
	if q == nil {
		writeError(w, 404, "unknown quiz")
		return
	}
	s.startGroup(w, q)
}
func (s *Server) uploadGroup(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, quiz.MaxFileSize)
	f, _, err := r.FormFile("file")
	if r.MultipartForm != nil {
		defer r.MultipartForm.RemoveAll()
	}
	if err != nil {
		writeError(w, 400, "expected a quiz file (maximum 8 MiB)")
		return
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		writeError(w, 400, "cannot read quiz")
		return
	}
	q, err := quiz.Parse(data)
	if err != nil {
		writeError(w, 422, err.Error())
		return
	}
	s.startGroup(w, q)
}
