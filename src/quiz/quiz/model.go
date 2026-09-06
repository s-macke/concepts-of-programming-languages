// Package quiz contains the domain model for quizzes: parsing YAML quiz
// files, sanitizing them for delivery to the browser and grading answers.
//
// The central design rule is that the correct answers never leave the server.
// Question.Sanitize strips them, and grading happens in Grade below.
package quiz

import (
	"fmt"
	"strings"
)

// Type enumerates the supported question types.
type Type string

const (
	// TypeSingle has exactly one correct option (radio buttons).
	TypeSingle Type = "single"
	// TypeMultiple has one or more correct options (checkboxes).
	TypeMultiple Type = "multiple"
	// TypeTrueFalse is sugar for a single choice between true and false.
	TypeTrueFalse Type = "truefalse"
	// TypeText is a free text answer compared against accepted strings.
	TypeText Type = "text"
)

// Quiz is a complete quiz as read from a YAML file.
type Quiz struct {
	Title       string     `yaml:"title"`
	Description string     `yaml:"description"`
	Questions   []Question `yaml:"questions"`
}

// Image is an inline, base64 encoded image. Quiz files are self-contained so
// that a single file can be uploaded without any accompanying assets.
type Image struct {
	Data string `yaml:"data" json:"data"`
	MIME string `yaml:"mime" json:"mime,omitempty"`
	Alt  string `yaml:"alt" json:"alt,omitempty"`
}

// Code is an optional source code snippet shown with a question.
type Code struct {
	Lang   string `yaml:"lang" json:"lang,omitempty"`
	Source string `yaml:"source" json:"source"`
}

// Option is one selectable answer of a choice question.
type Option struct {
	ID    string `yaml:"id" json:"id"`
	Text  string `yaml:"text" json:"text"`
	Image *Image `yaml:"image" json:"image,omitempty"`
}

// Question is a single question of a quiz. Answer holds the correct
// answer(s) and is never serialized towards the client.
type Question struct {
	Type        Type     `yaml:"type"`
	Text        string   `yaml:"text"`
	Code        *Code    `yaml:"code"`
	Image       *Image   `yaml:"image"`
	Options     []Option `yaml:"options"`
	Answer      Answer   `yaml:"answer"`
	Explanation string   `yaml:"explanation"`
	Points      int      `yaml:"points"`
}

// PublicQuestion is the client facing view of a Question: same content, but
// without Answer and Explanation.
type PublicQuestion struct {
	Type    Type     `json:"type"`
	Text    string   `json:"text"`
	Code    *Code    `json:"code,omitempty"`
	Image   *Image   `json:"image,omitempty"`
	Options []Option `json:"options,omitempty"`
	Points  int      `json:"points"`
}

// Sanitize returns the question without the information a student must not see.
func (q *Question) Sanitize() PublicQuestion {
	return PublicQuestion{
		Type:    q.Type,
		Text:    q.Text,
		Code:    q.Code,
		Image:   q.Image,
		Options: q.Options,
		Points:  q.Points,
	}
}

// Answer holds the accepted answers of a question. YAML allows both a scalar
// ("answer: a") and a list ("answer: [a, c]"); both end up here as a slice.
type Answer []string

// UnmarshalYAML accepts either a scalar or a sequence, so that single choice
// questions can use the more natural scalar form.
func (a *Answer) UnmarshalYAML(unmarshal func(any) error) error {
	var scalar string
	if err := unmarshal(&scalar); err == nil {
		*a = Answer{scalar}
		return nil
	}
	var list []string
	if err := unmarshal(&list); err != nil {
		return fmt.Errorf("answer must be a string or a list of strings")
	}
	*a = list
	return nil
}

// normalize prepares a free text answer for comparison.
func normalize(s string) string {
	return strings.Join(strings.Fields(strings.ToLower(s)), " ")
}
