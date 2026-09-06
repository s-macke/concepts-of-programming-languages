package quiz

// Result is the feedback a student receives after answering one question.
type Result struct {
	Correct       bool     `json:"correct"`
	Points        int      `json:"points"`
	CorrectAnswer []string `json:"correctAnswer"`
	Explanation   string   `json:"explanation,omitempty"`
}

// Grade compares the given answer against the correct one. The answer is
// always a slice: choice questions send option ids, text questions send a
// single string.
//
// Multiple choice is graded all or nothing: the selected set must match the
// correct set exactly.
func (q *Question) Grade(given []string) Result {
	correct := false

	switch q.Type {
	case TypeText:
		correct = len(given) == 1 && q.matchesText(given[0])
	case TypeMultiple:
		correct = sameSet(given, q.Answer)
	default: // single, truefalse
		correct = len(given) == 1 && given[0] == q.Answer[0]
	}

	res := Result{
		Correct:       correct,
		CorrectAnswer: q.Answer,
		Explanation:   q.Explanation,
	}
	if correct {
		res.Points = q.Points
	}
	return res
}

// matchesText reports whether s matches one of the accepted answers, ignoring
// case and surrounding or repeated whitespace.
func (q *Question) matchesText(s string) bool {
	got := normalize(s)
	for _, want := range q.Answer {
		if got == normalize(want) {
			return true
		}
	}
	return false
}

// sameSet reports whether a and b contain the same elements, ignoring order
// and duplicates.
func sameSet(a, b []string) bool {
	set := make(map[string]bool, len(b))
	for _, s := range b {
		set[s] = true
	}
	seen := make(map[string]bool, len(a))
	for _, s := range a {
		if !set[s] {
			return false
		}
		seen[s] = true
	}
	return len(seen) == len(set)
}
