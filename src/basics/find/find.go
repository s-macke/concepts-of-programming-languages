package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func ListDir(path string, pattern string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		name := path + string(os.PathSeparator) + f.Name()
		b, err := regexp.MatchString(pattern, f.Name())
		if err != nil {

			log.Fatal(err)
		}
		if b {
			fmt.Println(name)
		}
		if f.IsDir() {
			ListDir(name, pattern)
		}
	}
}

func main() {
	path := flag.String("path", ".", "path to search")
	regex := flag.String("regex", ".*", "path")
	flag.Parse()
	ListDir(*path, *regex)
}
