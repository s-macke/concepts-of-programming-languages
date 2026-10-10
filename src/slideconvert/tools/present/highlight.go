package present

import (
	"fmt"
	stdhtml "html"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/alecthomas/chroma/v2"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
)

// highlightCode adds syntax-highlighted HTML to source lines. The file name is
// the primary language hint; extensionless files fall back to content analysis.
// If Chroma cannot identify or tokenize the language, the caller renders the
// original escaped text instead.
func highlightCode(lines []codeLine, filename string) ([]codeLine, bool) {
	if len(lines) == 0 {
		return lines, false
	}

	sourceLines := make([]string, len(lines))
	for i, line := range lines {
		sourceLines[i] = line.L
	}
	source := strings.Join(sourceLines, "\n")

	lexer := lexers.Match(filename)
	if isFallbackLexer(lexer) {
		lexer = lexers.Analyse(source)
	}
	if isFallbackLexer(lexer) && filepath.Ext(filename) == "" && looksLikeGo(source) {
		lexer = lexers.Get("go")
	}
	if isFallbackLexer(lexer) {
		return lines, false
	}
	return highlightWithLexer(lines, lexer)
}

// highlightPre highlights an inline preformatted block from a .slide file.
// The language is given explicitly by a preceding "#lang" comment.
func highlightPre(source, lang string) (template.HTML, bool) {
	lexer := lexers.Get(lang)
	if isFallbackLexer(lexer) {
		return "", false
	}
	sourceLines := strings.Split(source, "\n")
	lines := make([]codeLine, len(sourceLines))
	for i, l := range sourceLines {
		lines[i] = codeLine{L: l, N: i + 1}
	}
	lines, ok := highlightWithLexer(lines, lexer)
	if !ok {
		return "", false
	}
	rendered := make([]string, len(lines))
	for i, l := range lines {
		rendered[i] = string(l.HTML)
	}
	return template.HTML(strings.Join(rendered, "\n")), true
}

// highlightWithLexer renders each source line as HTML with inline styles.
func highlightWithLexer(lines []codeLine, lexer chroma.Lexer) ([]codeLine, bool) {
	sourceLines := make([]string, len(lines))
	for i, line := range lines {
		sourceLines[i] = line.L
	}
	source := strings.Join(sourceLines, "\n")

	iterator, err := chroma.Coalesce(lexer).Tokenise(nil, source)
	if err != nil {
		return lines, false
	}
	tokenLines := chroma.SplitTokensIntoLines(iterator.Tokens())
	if len(tokenLines) != len(lines) {
		return lines, false
	}

	theme := styles.GitHub
	background := theme.Get(chroma.Background)
	for i, tokens := range tokenLines {
		var rendered strings.Builder
		leadingToTrim := 0
		if lines[i].HL {
			leadingToTrim = len(leadingSpaceRE.FindString(lines[i].L))
		}
		for _, token := range tokens {
			value := strings.TrimSuffix(token.Value, "\n")
			if leadingToTrim >= len(value) {
				leadingToTrim -= len(value)
				continue
			}
			if leadingToTrim > 0 {
				value = value[leadingToTrim:]
				leadingToTrim = 0
			}
			if value == "" {
				continue
			}
			escaped := stdhtml.EscapeString(value)
			if strings.TrimSpace(value) == "" {
				rendered.WriteString(escaped)
				continue
			}
			css := chromahtml.StyleEntryToCSS(theme.Get(token.Type).Sub(background))
			if css == "" {
				rendered.WriteString(escaped)
				continue
			}
			fmt.Fprintf(&rendered, `<span style="%s">%s</span>`, css, escaped)
		}
		lines[i].HTML = template.HTML(rendered.String())
	}
	return lines, true
}

func isFallbackLexer(lexer chroma.Lexer) bool {
	return lexer == nil || lexer.Config().Name == lexers.Fallback.Config().Name
}

func looksLikeGo(source string) bool {
	for _, marker := range []string{"go func(", "make(chan ", "<-", ":="} {
		if strings.Contains(source, marker) {
			return true
		}
	}
	return false
}
