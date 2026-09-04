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
