package quiz

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// MaxFileSize is the largest quiz file that is accepted, both from disk and
// from an upload. Base64 images make quiz files big, but not arbitrarily big.
const MaxFileSize = 8 << 20 // 8 MiB

// Parse reads a quiz from YAML and validates it.
func Parse(data []byte) (*Quiz, error) {
	var q Quiz
	if err := yaml.Unmarshal(data, &q); err != nil {
		return nil, fmt.Errorf("invalid YAML: %w", err)
	}
	if err := q.validate(); err != nil {
		return nil, err
	}
	return &q, nil
}

// ParseFile reads and validates a quiz file from disk.
func ParseFile(path string) (*Quiz, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if info.Size() > MaxFileSize {
		return nil, fmt.Errorf("%s is larger than %d bytes", path, MaxFileSize)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	q, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", filepath.Base(path), err)
	}
	return q, nil
}

// validate checks the invariants the grader and the UI rely on. Doing this
// once at load time keeps every later stage free of defensive checks.
func (q *Quiz) validate() error {
	if strings.TrimSpace(q.Title) == "" {
		return fmt.Errorf("quiz needs a title")
	}
	if len(q.Questions) == 0 {
		return fmt.Errorf("quiz %q has no questions", q.Title)
	}
	for i := range q.Questions {
		if err := q.Questions[i].validate(); err != nil {
			return fmt.Errorf("question %d: %w", i+1, err)
		}
	}
	return nil
}

// trueFalseOptions are used when a truefalse question omits its options.
var trueFalseOptions = []Option{{ID: "true", Text: "True"}, {ID: "false", Text: "False"}}

func (q *Question) validate() error {
	if q.Points == 0 {
		q.Points = 1
	}
	if q.Points < 0 {
		return fmt.Errorf("points must not be negative")
	}
	if strings.TrimSpace(q.Text) == "" {
		return fmt.Errorf("needs a text")
	}
	if len(q.Answer) == 0 {
		return fmt.Errorf("needs an answer")
	}
	if q.Type == "" {
		q.Type = TypeSingle
	}

	switch q.Type {
	case TypeTrueFalse:
		if len(q.Options) == 0 {
			q.Options = trueFalseOptions
		}
		fallthrough
	case TypeSingle:
		if len(q.Answer) != 1 {
			return fmt.Errorf("type %s takes exactly one answer", q.Type)
		}
		return q.validateOptions()
	case TypeMultiple:
		return q.validateOptions()
	case TypeText:
		if len(q.Options) > 0 {
			return fmt.Errorf("type text must not have options")
		}
		return nil
	default:
		return fmt.Errorf("unknown type %q", q.Type)
	}
}

// validateOptions makes sure the options are usable and that every declared
// answer actually refers to one of them.
func (q *Question) validateOptions() error {
	if len(q.Options) < 2 {
		return fmt.Errorf("needs at least two options")
	}
	ids := make(map[string]bool, len(q.Options))
	for i, o := range q.Options {
		if o.ID == "" {
			return fmt.Errorf("option %d needs an id", i+1)
		}
		if ids[o.ID] {
			return fmt.Errorf("duplicate option id %q", o.ID)
		}
		ids[o.ID] = true
	}
	for _, a := range q.Answer {
		if !ids[a] {
			return fmt.Errorf("answer %q is not one of the options", a)
		}
	}
	return nil
}

// TotalPoints is the score a perfect run achieves.
func (q *Quiz) TotalPoints() int {
	total := 0
	for _, question := range q.Questions {
		total += question.Points
	}
	return total
}

// LoadDir reads every *.yaml and *.yml file of dir. The returned map is keyed
// by the file name without extension, which is also the quiz id used by the
// HTTP API. Broken files are reported but do not prevent the server from
// starting with the remaining quizzes.
func LoadDir(dir string) (map[string]*Quiz, []error) {
	quizzes := make(map[string]*Quiz)
	var problems []error

	entries, err := os.ReadDir(dir)
	if err != nil {
		return quizzes, []error{err}
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(e.Name()))
		if ext != ".yaml" && ext != ".yml" {
			continue
		}
		q, err := ParseFile(filepath.Join(dir, e.Name()))
		if err != nil {
			problems = append(problems, err)
			continue
		}
		quizzes[strings.TrimSuffix(e.Name(), filepath.Ext(e.Name()))] = q
	}
	return quizzes, problems
}

// SortedIDs returns the quiz ids in a stable order for listing.
func SortedIDs(quizzes map[string]*Quiz) []string {
	ids := make([]string, 0, len(quizzes))
	for id := range quizzes {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}
