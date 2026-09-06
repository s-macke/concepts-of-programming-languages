package present

import (
	"html/template"
	"strings"
	"testing"
)

func TestHighlightCode(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		source   []string
		wantText string
	}{
		{
			name:     "Go",
			filename: "main.go",
			source:   []string{"package main", "func main() {}"},
			wantText: "color: #cf222e",
		},
		{
			name:     "C",
			filename: "main.c",
			source:   []string{"int main(void) {", "}"},
			wantText: "color:",
		},
		{
			name:     "Java",
			filename: "Main.java",
			source:   []string{"class Main {}"},
			wantText: "color:",
		},
		{
			name:     "JavaScript",
			filename: "main.js",
			source:   []string{"const answer = 42;"},
			wantText: "color:",
		},
		{
			name:     "HTML",
			filename: "index.html",
			source:   []string{"<main>Hello</main>"},
			wantText: "&lt;",
		},
		{
			name:     "WebAssembly text",
			filename: "add.wat",
			source:   []string{"(module (func))"},
			wantText: "color:",
		},
		{
			name:     "Makefile",
			filename: "Makefile",
			source:   []string{"build:", "\tgo build ./..."},
			wantText: "color:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines := makeCodeLines(tt.source)
			got, ok := highlightCode(lines, tt.filename)
			if !ok {
				t.Fatalf("expected %s source to be highlighted", tt.name)
			}
			if rendered := joinHighlightedLines(got); !strings.Contains(rendered, tt.wantText) {
				t.Errorf("highlighted HTML %q does not contain %q", rendered, tt.wantText)
			}
		})
	}
}

func TestHighlightCodeDetectsExtensionlessGo(t *testing.T) {
	lines := makeCodeLines([]string{
		`f("hello", "world")`,
		`go f("hello", "world")`,
		"timerChan := make(chan time.Time)",
		"go func() { timerChan <- time.Now() }()",
	})

	got, ok := highlightCode(lines, "snippet")
	if !ok {
		t.Fatal("expected extensionless Go source to be detected")
	}
	if rendered := joinHighlightedLines(got); !strings.Contains(rendered, "color: #cf222e") {
		t.Errorf("expected Go keyword highlighting, got %q", rendered)
	}
}

func TestHighlightCodeEscapesSource(t *testing.T) {
	lines := makeCodeLines([]string{"package main", "func less(a, b int) bool { return a < b }"})

	got, ok := highlightCode(lines, "main.go")
	if !ok {
		t.Fatal("expected Go source to be highlighted")
	}
	rendered := joinHighlightedLines(got)
	if strings.Contains(rendered, "a < b") || !strings.Contains(rendered, "&lt;") {
		t.Errorf("source was not HTML-escaped: %q", rendered)
	}
}

func TestHighlightCodeFallsBackForUnknownText(t *testing.T) {
	lines := makeCodeLines([]string{"ordinary unstructured words"})

	got, ok := highlightCode(lines, "notes.unknown-extension")
	if ok {
		t.Fatalf("unexpected highlighting for unknown text: %q", joinHighlightedLines(got))
	}
	if got[0].HTML != template.HTML("") {
		t.Errorf("fallback should leave highlighted HTML empty, got %q", got[0].HTML)
	}
}

func makeCodeLines(source []string) []codeLine {
	lines := make([]codeLine, len(source))
	for i, line := range source {
		lines[i] = codeLine{L: line, N: i + 1}
	}
	return lines
}

func joinHighlightedLines(lines []codeLine) string {
	values := make([]string, len(lines))
	for i, line := range lines {
		values[i] = string(line.HTML)
	}
	return strings.Join(values, "\n")
}
