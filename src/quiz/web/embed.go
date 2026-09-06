// Package web embeds the compiled TypeScript frontend into the server binary,
// so that a single `go run ./src/quiz/cmd/quizserver` serves the whole
// application without any Node.js toolchain present.
//
// The contents of dist/ are checked in. Rebuild them with `npm run build` in
// this directory after changing anything under src/.
package web

import (
	"embed"
	"io/fs"
)

//go:embed all:dist
var dist embed.FS

// Assets returns the built frontend rooted at dist/.
func Assets() fs.FS {
	sub, err := fs.Sub(dist, "dist")
	if err != nil {
		// Can only happen if the embed directive above is broken.
		panic(err)
	}
	return sub
}
