# slideconvert

the slides of this lecture are created via the Go
[present tool](https://pkg.go.dev/golang.org/x/tools/present)
and viewed as a web page.

Unfortunately, you can't view them offline with this tool.

This code here is an extraction of the rendering part of the .slide
files and creates the .html slides linked in the main README.

Source included with `.code` and `.play` is syntax-highlighted during
rendering. The language is selected from the source filename, with content
detection as a fallback, so the generated slides remain static and work
without client-side JavaScript.

Indented code blocks written directly in a .slide file are highlighted only
when a `#lang` line precedes them. For the original present tool this line is
an ordinary comment, so `make slideshow` keeps working. The hint is invisible
and applies to the next code block only. Any Chroma language name works, e.g. `go`, `c`,
`java`, `python` or `javascript`:

    #lang python

        def greet(name: str) -> str:
            return "Hello " + name
