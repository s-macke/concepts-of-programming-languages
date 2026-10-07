package group

import (
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

const fixture = `title: Group test
questions:
- text: Pick A
  options: [{id: a, text: A}, {id: b, text: B}]
  answer: a
  explanation: Secret explanation
- type: multiple
  text: Pick both
  options: [{id: a, text: A}, {id: b, text: B}]
  answer: [a, b]
- type: text
  text: Keyword
  answer: go
`

func setup(t *testing.T) (*Store, string, string) {
	t.Helper()
	q, e := quiz.Parse([]byte(fixture))
	if e != nil {
		t.Fatal(e)
	}
	s := NewStore(time.Hour)
	id, h := s.Create(q)
	return s, id, h
}
func must(t *testing.T, e error) {
	t.Helper()
	if e != nil {
		t.Fatal(e)
	}
}
func TestLifecyclePrivacyAndStatistics(t *testing.T) {
	s, id, h := setup(t)
	a, e := s.Join(id)
	must(t, e)
	b, e := s.Join(id)
	must(t, e)
	if _, e = s.State(id, "invalid"); !errors.Is(e, ErrAuth) {
		t.Fatal(e)
	}
	if e = s.Action(id, a, "start", 0); !errors.Is(e, ErrAuth) {
		t.Fatal(e)
	}
	if e = s.Answer(id, a, 1, []string{"a"}); !errors.Is(e, ErrConflict) {
		t.Fatal(e)
	}
	must(t, s.Action(id, h, "start", 0))
	if e = s.Action(id, h, "start", 0); !errors.Is(e, ErrConflict) {
		t.Fatal(e)
	}
	if e = s.Answer(id, a, 1, []string{"invalid"}); !errors.Is(e, ErrAnswer) {
		t.Fatal(e)
	}
	must(t, s.Answer(id, a, 1, []string{"a"}))
	must(t, s.Answer(id, a, 1, []string{"b"}))
	must(t, s.Answer(id, b, 1, []string{"b"}))
	for _, tok := range []string{h, a, b} {
		st, e := s.State(id, tok)
		must(t, e)
		data, _ := json.Marshal(st)
		for _, secret := range []string{"correctAnswer", "Secret explanation", "personal", "stats"} {
			if strings.Contains(string(data), secret) {
				t.Fatalf("leaked %s: %s", secret, data)
			}
		}
		if st.Submitted != 2 {
			t.Fatal(st)
		}
	}
	must(t, s.Action(id, h, "reveal", 1))
	st, e := s.State(id, a)
	must(t, e)
	if st.Stats.Correct != 1 || st.Stats.Incorrect != 1 || st.Stats.Selections["a"] != 1 || st.Personal == nil || !st.Personal.Correct {
		t.Fatalf("%+v", st)
	}
	late, e := s.Join(id)
	must(t, e)
	st, e = s.State(id, late)
	must(t, e)
	if st.Stats.Unanswered != 1 {
		t.Fatal(st.Stats)
	}
	if e = s.Answer(id, late, 1, []string{"a"}); !errors.Is(e, ErrConflict) {
		t.Fatal(e)
	}
	must(t, s.Action(id, h, "next", 2))
	must(t, s.Answer(id, a, 2, []string{"a", "b"}))
	must(t, s.Action(id, h, "reveal", 3))
	st, e = s.State(id, h)
	must(t, e)
	if st.Personal != nil || st.Stats.Selections["a"] != 1 || st.Stats.Selections["b"] != 1 {
		t.Fatal(st)
	}
	must(t, s.Action(id, h, "next", 4))
	must(t, s.Answer(id, a, 3, []string{" GO "}))
	must(t, s.Answer(id, b, 3, []string{"Private name must not be kept"}))
	must(t, s.Action(id, h, "finish", 5))
	st, e = s.State(id, h)
	must(t, e)
	if st.Summary.Active != 2 || st.Summary.Completed != 1 || st.Summary.Submissions != 5 || st.Summary.Average != 50 {
		t.Fatalf("%+v", st.Summary)
	}
	data, _ := json.Marshal(s.rooms[id])
	if strings.Contains(string(data), "Private name") {
		t.Fatal("raw text retained")
	}
	if _, e = s.Join(id); !errors.Is(e, ErrConflict) {
		t.Fatal(e)
	}
	must(t, s.Action(id, h, "delete", 6))
	if _, e = s.State(id, h); !errors.Is(e, ErrMissing) {
		t.Fatal(e)
	}
}
func TestConcurrentAnswersAndReveal(t *testing.T) {
	s, id, h := setup(t)
	participants := []string{}
	for i := 0; i < 80; i++ {
		p, e := s.Join(id)
		must(t, e)
		participants = append(participants, p)
	}
	must(t, s.Action(id, h, "start", 0))
	var wg sync.WaitGroup
	for _, p := range participants {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				e := s.Answer(id, p, 1, []string{"a"})
				if e != nil && !errors.Is(e, ErrConflict) {
					t.Error(e)
				}
				s.State(id, p)
			}
		}(p)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if e := s.Action(id, h, "reveal", 1); e != nil {
			t.Error(e)
		}
	}()
	wg.Wait()
	st, e := s.State(id, h)
	must(t, e)
	if st.Stats.Submitted > 80 || st.Stats.Correct != st.Stats.Submitted || st.Stats.Selections["a"] != st.Stats.Submitted || st.Stats.Unanswered+st.Stats.Submitted != 80 {
		t.Fatal(st.Stats)
	}
}
func TestExpiryAndEarlyFinish(t *testing.T) {
	s, id, h := setup(t)
	s.rooms[id].LastSeen = time.Now().Add(-2 * time.Hour)
	s.Expire()
	if _, e := s.State(id, h); !errors.Is(e, ErrMissing) {
		t.Fatal(e)
	}
	s, id, h = setup(t)
	must(t, s.Action(id, h, "finish", 0))
	st, e := s.State(id, h)
	must(t, e)
	if st.Summary.Average != 0 || len(st.Summary.Questions) != 0 {
		t.Fatal(st)
	}
}
