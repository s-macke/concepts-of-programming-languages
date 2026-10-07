// Package group implements anonymous, host-led quiz rooms in memory.
package group

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

var ErrMissing = errors.New("unknown or expired room")
var ErrAuth = errors.New("invalid room credential")
var ErrConflict = errors.New("room has changed or action is unavailable")
var ErrAnswer = errors.New("invalid answer")

func token() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

type Participant struct{ Results map[int]quiz.Result }
type Stats struct {
	Number     int            `json:"number"`
	Text       string         `json:"text"`
	Submitted  int            `json:"submitted"`
	Unanswered int            `json:"unanswered"`
	Correct    int            `json:"correct"`
	Incorrect  int            `json:"incorrect"`
	Selections map[string]int `json:"selections,omitempty"`
}
type Summary struct {
	Active      int     `json:"active"`
	Completed   int     `json:"completed"`
	Submissions int     `json:"submissions"`
	Average     float64 `json:"average"`
	Questions   []Stats `json:"questions"`
}
type State struct {
	Title        string               `json:"title"`
	Phase        string               `json:"phase"`
	Revision     int                  `json:"revision"`
	Number       int                  `json:"number"`
	Total        int                  `json:"total"`
	Participants int                  `json:"participants"`
	Submitted    int                  `json:"submitted"`
	Answered     bool                 `json:"answered"`
	Question     *quiz.PublicQuestion `json:"question,omitempty"`
	Answer       []string             `json:"correctAnswer,omitempty"`
	Explanation  string               `json:"explanation,omitempty"`
	Personal     *quiz.Result         `json:"personal,omitempty"`
	Stats        *Stats               `json:"stats,omitempty"`
	Summary      *Summary             `json:"summary,omitempty"`
}
type Room struct {
	ID, Host        string
	Quiz            *quiz.Quiz
	Phase           string
	Revision, Index int
	Participants    map[string]*Participant
	Stats           []Stats
	LastSeen        time.Time
}

// Store serializes room transitions, submissions, snapshots and expiry together.
type Store struct {
	mu    sync.Mutex
	rooms map[string]*Room
	ttl   time.Duration
}

func NewStore(ttl time.Duration) *Store { return &Store{rooms: map[string]*Room{}, ttl: ttl} }
func (s *Store) Create(q *quiz.Quiz) (string, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := &Room{ID: token(), Host: token(), Quiz: q, Phase: "lobby", Index: -1, Participants: map[string]*Participant{}, LastSeen: time.Now()}
	for i, q := range q.Questions {
		r.Stats = append(r.Stats, Stats{Number: i + 1, Text: q.Text, Selections: map[string]int{}})
	}
	s.rooms[r.ID] = r
	return r.ID, r.Host
}
func (s *Store) get(id string) (*Room, error) {
	r := s.rooms[id]
	if r == nil {
		return nil, ErrMissing
	}
	if time.Since(r.LastSeen) > s.ttl {
		delete(s.rooms, id)
		return nil, ErrMissing
	}
	return r, nil
}
func (s *Store) Expire() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for id := range s.rooms {
		s.get(id)
	}
}
func (s *Store) Join(id string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, err := s.get(id)
	if err != nil {
		return "", err
	}
	if r.Phase == "finished" {
		return "", ErrConflict
	}
	t := token()
	r.Participants[t] = &Participant{Results: map[int]quiz.Result{}}
	r.LastSeen = time.Now()
	return t, nil
}
func (r *Room) auth(t string) bool { return t == r.Host || r.Participants[t] != nil }
func (s *Store) Action(id, t, action string, revision int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, err := s.get(id)
	if err != nil {
		return err
	}
	if t != r.Host {
		return ErrAuth
	}
	if action == "delete" {
		delete(s.rooms, id)
		return nil
	}
	if revision != r.Revision {
		return ErrConflict
	}
	switch {
	case action == "start" && r.Phase == "lobby":
		r.Index = 0
		r.Phase = "question"
	case action == "reveal" && r.Phase == "question":
		r.Phase = "reveal"
	case action == "next" && r.Phase == "reveal" && r.Index+1 < len(r.Quiz.Questions):
		r.Index++
		r.Phase = "question"
	case action == "finish" && r.Phase != "finished":
		r.Phase = "finished"
	default:
		return ErrConflict
	}
	r.Revision++
	r.LastSeen = time.Now()
	return nil
}
func (s *Store) Answer(id, t string, n int, given []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, err := s.get(id)
	if err != nil {
		return err
	}
	p := r.Participants[t]
	if p == nil {
		return ErrAuth
	}
	if _, ok := p.Results[n-1]; ok {
		r.LastSeen = time.Now()
		return nil
	}
	if r.Phase != "question" || n-1 != r.Index {
		return ErrConflict
	}
	q := r.Quiz.Questions[r.Index]
	if len(given) == 0 || (q.Type != quiz.TypeMultiple && len(given) != 1) {
		return ErrAnswer
	}
	seen := map[string]bool{}
	for _, a := range given {
		if q.Type == quiz.TypeText {
			if strings.TrimSpace(a) == "" {
				return ErrAnswer
			}
			continue
		}
		valid := false
		for _, o := range q.Options {
			if o.ID == a {
				valid = true
			}
		}
		if !valid || seen[a] {
			return ErrAnswer
		}
		seen[a] = true
	}
	res := q.Grade(given)
	p.Results[r.Index] = res // Raw text is deliberately never retained.
	st := &r.Stats[r.Index]
	st.Submitted++
	if res.Correct {
		st.Correct++
	} else {
		st.Incorrect++
	}
	for a := range seen {
		st.Selections[a]++
	}
	r.LastSeen = time.Now()
	return nil
}
func (s *Store) State(id, t string) (State, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	r, err := s.get(id)
	if err != nil {
		return State{}, err
	}
	if !r.auth(t) {
		return State{}, ErrAuth
	}
	r.LastSeen = time.Now()
	st := State{Title: r.Quiz.Title, Phase: r.Phase, Revision: r.Revision, Number: r.Index + 1, Total: len(r.Quiz.Questions), Participants: len(r.Participants)}
	copyStats := func(i int) Stats {
		v := r.Stats[i]
		v.Unanswered = len(r.Participants) - v.Submitted
		v.Selections = map[string]int{}
		for k, n := range r.Stats[i].Selections {
			v.Selections[k] = n
		}
		return v
	}
	if r.Index >= 0 {
		q := r.Quiz.Questions[r.Index]
		pub := q.Sanitize()
		st.Question = &pub
		st.Submitted = r.Stats[r.Index].Submitted
		if p := r.Participants[t]; p != nil {
			res, ok := p.Results[r.Index]
			st.Answered = ok
			if ok && r.Phase != "question" {
				st.Personal = &res
			}
		}
		if r.Phase == "reveal" || r.Phase == "finished" {
			st.Answer = q.Answer
			st.Explanation = q.Explanation
			v := copyStats(r.Index)
			st.Stats = &v
		}
	}
	if r.Phase == "finished" {
		sum := &Summary{Questions: []Stats{}}
		points := 0
		for _, p := range r.Participants {
			if len(p.Results) > 0 {
				sum.Active++
			}
			if len(p.Results) == len(r.Quiz.Questions) {
				sum.Completed++
			}
			sum.Submissions += len(p.Results)
			for _, v := range p.Results {
				points += v.Points
			}
		}
		if sum.Active > 0 && r.Quiz.TotalPoints() > 0 {
			sum.Average = 100 * float64(points) / float64(sum.Active*r.Quiz.TotalPoints())
		}
		for i := 0; i <= r.Index; i++ {
			sum.Questions = append(sum.Questions, copyStats(i))
		}
		st.Summary = sum
	}
	return st, nil
}
