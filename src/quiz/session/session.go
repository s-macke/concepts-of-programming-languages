// Package session tracks the state of a running quiz.
//
// The Session interface is the seam that keeps a future live, classroom wide
// mode cheap to add: question delivery and grading are defined here, so a
// hosted session with many participants and a WebSocket hub can be added
// beside SoloSession without touching the HTTP API or the domain model.
package session

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

// Session is one run of a quiz.
type Session interface {
	ID() string
	Info() Info
	// Question returns the sanitized question with the given index.
	Question(index int) (quiz.PublicQuestion, error)
	// Answer grades the given answer and records the score.
	Answer(index int, given []string) (quiz.Result, error)
	// Summary reports the score achieved so far.
	Summary() Summary
	// LastSeen is used by the store to expire idle sessions.
	LastSeen() time.Time
	touch()
}

// Info describes a session to the client without revealing any question.
type Info struct {
	SessionID    string `json:"sessionId"`
	Title        string `json:"title"`
	Description  string `json:"description,omitempty"`
	NumQuestions int    `json:"numQuestions"`
	TotalPoints  int    `json:"totalPoints"`
}

// Summary is the final (or intermediate) score of a session.
type Summary struct {
	Answered     int `json:"answered"`
	NumQuestions int `json:"numQuestions"`
	Points       int `json:"points"`
	TotalPoints  int `json:"totalPoints"`
	Correct      int `json:"correct"`
}

// newID returns a random, unguessable session id.
func newID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", fmt.Errorf("cannot generate session id: %w", err)
	}
	return hex.EncodeToString(b), nil
}
