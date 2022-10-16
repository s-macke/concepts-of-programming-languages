// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

var (
	basePath    = flag.String("base", "", "base path for slide template and static resources")
	contentPath = flag.String("content", ".", "base path for presentation content")
)

func main() {
	flag.Parse()

	err := initTemplates(*basePath)
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	files, err := filepath.Glob(*contentPath + string(os.PathSeparator) + "*.slide")
	if err != nil {
		log.Fatal(err)
	}
	for _, filename := range files {
		fmt.Println(filename)
		buf := new(bytes.Buffer)
		err = renderDoc(buf, filename)
		if err != nil {
			log.Fatalf("Failed to render doc: %v", err)
		}
		ext := path.Ext(filename)
		outfile := filename[0:len(filename)-len(ext)] + ".html"
		err = os.WriteFile(outfile, buf.Bytes(), 0644)
		//fmt.Println(buf)
	}
}
