// Command quizpack inlines images into a quiz file.
//
// The quiz format itself only knows self-contained, base64 encoded images, so
// that any single file can be uploaded. That is unpleasant to author and to
// review in a diff. quizpack lets you write
//
//	image: {path: diagrams/channels.png, alt: A channel}
//
// in a source file and turns it into the base64 form on the way out:
//
//	go run ./src/quiz/cmd/quizpack -in lecture7.src.yaml -out src/quiz/quizzes/lecture7.yaml
//
// Paths are resolved relative to the input file.
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/s-macke/concepts-of-programming-languages/src/quiz/quiz"
)

func main() {
	in := flag.String("in", "", "quiz file to read (with image path: entries)")
	out := flag.String("out", "", "file to write (default: stdout)")
	flag.Parse()

	if *in == "" {
		log.Fatal("missing -in")
	}
	data, err := os.ReadFile(*in)
	if err != nil {
		log.Fatal(err)
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		log.Fatalf("invalid YAML: %v", err)
	}
	baseDir := filepath.Dir(*in)
	inlined, err := inlineImages(&doc, baseDir)
	if err != nil {
		log.Fatal(err)
	}

	packed, err := yaml.Marshal(&doc)
	if err != nil {
		log.Fatal(err)
	}
	// Fail early rather than shipping a quiz the server will reject.
	if _, err := quiz.Parse(packed); err != nil {
		log.Fatalf("packed quiz is not valid: %v", err)
	}

	if *out == "" {
		os.Stdout.Write(packed)
	} else if err := os.WriteFile(*out, packed, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("inlined %d image(s)", inlined)
}

// inlineImages walks the document and replaces every "path" key of an image
// mapping with the "data" and "mime" keys the quiz format expects.
func inlineImages(node *yaml.Node, baseDir string) (int, error) {
	count := 0
	if node.Kind == yaml.MappingNode {
		for i := 0; i+1 < len(node.Content); i += 2 {
			key, value := node.Content[i], node.Content[i+1]
			if key.Value != "image" || value.Kind != yaml.MappingNode {
				continue
			}
			n, err := inlineImage(value, baseDir)
			if err != nil {
				return count, err
			}
			count += n
		}
	}
	for _, child := range node.Content {
		n, err := inlineImages(child, baseDir)
		if err != nil {
			return count, err
		}
		count += n
	}
	return count, nil
}

// inlineImage rewrites a single image mapping in place.
func inlineImage(image *yaml.Node, baseDir string) (int, error) {
	for i := 0; i+1 < len(image.Content); i += 2 {
		if image.Content[i].Value != "path" {
			continue
		}
		path := image.Content[i+1].Value
		raw, err := os.ReadFile(filepath.Join(baseDir, path))
		if err != nil {
			return 0, fmt.Errorf("image %s: %w", path, err)
		}
		mimeType := mime.TypeByExtension(strings.ToLower(filepath.Ext(path)))
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		// Replace the path key/value pair with data, and append mime.
		image.Content[i] = scalar("data")
		image.Content[i+1] = scalar(base64.StdEncoding.EncodeToString(raw))
		image.Content = append(image.Content, scalar("mime"), scalar(mimeType))
		return 1, nil
	}
	return 0, nil
}

func scalar(value string) *yaml.Node {
	return &yaml.Node{Kind: yaml.ScalarNode, Tag: "!!str", Value: value}
}
