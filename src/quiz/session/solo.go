package session

import (
	"fmt"
	"sync"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

// SoloSession is a single student working through a quiz alone. It is safe for
// concurrent use, because a browser may well have several requests in flight.
type SoloSession struct {
	id   string
	quiz *quiz.Quiz

	mu       sync.Mutex
	results  map[int]quiz.Result
	lastSeen time.Time
}

// NewSolo starts a new single player session for q.
func NewSolo(q *quiz.Quiz) (*SoloSession, error) {
	id, err := newID()
	if err != nil {
		return nil, err
	}
	return &SoloSession{
		id:       id,
		quiz:     q,
		results:  make(map[int]quiz.Result),
		lastSeen: time.Now(),
	}, nil
}

func (s *SoloSession) ID() string { return s.id }

func (s *SoloSession) Info() Info {
	return Info{
		SessionID:    s.id,
		Title:        s.quiz.Title,
		Description:  s.quiz.Description,
		NumQuestions: len(s.quiz.Questions),
		TotalPoints:  s.quiz.TotalPoints(),
	}
}

// ErrNoSuchQuestion is returned for an index outside the quiz.
var ErrNoSuchQuestion = fmt.Errorf("no such question")

func (s *SoloSession) Question(index int) (quiz.PublicQuestion, error) {
	if index < 0 || index >= len(s.quiz.Questions) {
		return quiz.PublicQuestion{}, ErrNoSuchQuestion
	}
	s.touch()
	return s.quiz.Questions[index].Sanitize(), nil
}

// Answer grades an answer. Answering the same question twice returns the
// original result, so that a reload cannot be used to try again.
func (s *SoloSession) Answer(index int, given []string) (quiz.Result, error) {
	if index < 0 || index >= len(s.quiz.Questions) {
		return quiz.Result{}, ErrNoSuchQuestion
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastSeen = time.Now()

	if res, ok := s.results[index]; ok {
		return res, nil
	}
	res := s.quiz.Questions[index].Grade(given)
	s.results[index] = res
	return res, nil
}

func (s *SoloSession) Summary() Summary {
	s.mu.Lock()
	defer s.mu.Unlock()

	sum := Summary{
		NumQuestions: len(s.quiz.Questions),
		TotalPoints:  s.quiz.TotalPoints(),
		Answered:     len(s.results),
	}
	for _, r := range s.results {
		sum.Points += r.Points
		if r.Correct {
			sum.Correct++
		}
	}
	return sum
}

func (s *SoloSession) LastSeen() time.Time {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.lastSeen
}

func (s *SoloSession) touch() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastSeen = time.Now()
}
