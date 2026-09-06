package session

import (
	"testing"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

func testQuiz(t *testing.T) *quiz.Quiz {
	t.Helper()
	q, err := quiz.Parse([]byte(`
title: T
questions:
  - type: single
    text: Pick a
    options: [{id: a, text: A}, {id: b, text: B}]
    answer: a
`))
	if err != nil {
		t.Fatal(err)
	}
	return q
}

func TestSoloSessionIsAnsweredOnlyOnce(t *testing.T) {
	s, err := NewSolo(testQuiz(t))
	if err != nil {
		t.Fatal(err)
	}
	if res, _ := s.Answer(0, []string{"b"}); res.Correct {
		t.Fatal("wrong answer graded as correct")
	}
	// Retrying must return the first, recorded result.
	res, _ := s.Answer(0, []string{"a"})
	if res.Correct {
		t.Error("a second attempt was accepted")
	}
	if sum := s.Summary(); sum.Points != 0 || sum.Answered != 1 {
		t.Errorf("summary = %+v", sum)
	}
}

func TestSoloSessionRejectsUnknownQuestions(t *testing.T) {
	s, _ := NewSolo(testQuiz(t))
	if _, err := s.Question(5); err != ErrNoSuchQuestion {
		t.Errorf("Question(5) error = %v", err)
	}
	if _, err := s.Answer(-1, nil); err != ErrNoSuchQuestion {
		t.Errorf("Answer(-1) error = %v", err)
	}
}

func TestStoreEvictsIdleSessions(t *testing.T) {
	store := NewStore(time.Hour)
	s, _ := NewSolo(testQuiz(t))
	store.Add(s)

	if _, ok := store.Get(s.ID()); !ok {
		t.Fatal("session not found after Add")
	}
	// Nothing is idle yet.
	if n := store.evict(time.Now().Add(-time.Hour)); n != 0 {
		t.Errorf("evicted %d live sessions", n)
	}
	// Everything last used before "now" is idle.
	if n := store.evict(time.Now().Add(time.Minute)); n != 1 {
		t.Errorf("evicted %d sessions, want 1", n)
	}
	if store.Len() != 0 {
		t.Errorf("store still holds %d sessions", store.Len())
	}
}

// Session ids must be unguessable, since they are the only access control.
func TestSessionIDsAreUnique(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 100; i++ {
		s, err := NewSolo(testQuiz(t))
		if err != nil {
			t.Fatal(err)
		}
		if seen[s.ID()] {
			t.Fatalf("duplicate session id %s", s.ID())
		}
		seen[s.ID()] = true
	}
}
