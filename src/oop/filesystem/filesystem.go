package main

import (
	"embed"
	"fmt"
	"github.com/s-macke/concepts-of-programming-languages/src/oop/filesystem/mockfs"
	"io/fs"
	"log"
	"os"
)

func ShowAllFiles(fsys fs.FS) {
	path := "files"
	files, err := fs.ReadDir(fsys, path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		//bytes, err := fs.ReadFile(fsys, path + string(os.PathSeparator) + file.Name())
		bytes, err := fs.ReadFile(fsys, path+"/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: '%s'\n", file.Name(), string(bytes))
	}
}

//go:embed files
var fsEmbed embed.FS

func main() {

	fmt.Println("Show OS filesystem")
	fsOS := os.DirFS(".")
	ShowAllFiles(fsOS)

	fmt.Println("\nShow embedded filesystem")
	ShowAllFiles(fsEmbed)

	var mockfs = mockfs.NewMockFs("hello.txt", "Hello World")
	fmt.Println("\nShow mock filesystem")
	ShowAllFiles(mockfs)
}
