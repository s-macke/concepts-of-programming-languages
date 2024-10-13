package main

import (
	"fmt"
	"strings"
)

type Page []string

func (p Page) String() string {
	return strings.Join(p, "\n")
}

func main() {
	page := Page{"This is the first paragraph", "This is the second paragraph"}
	fmt.Println(page)
}
