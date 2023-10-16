package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello %s", "Programming with Go \xE2\x98\xAF\n") // \xE2\x98\xAF -> ☯
	fmt.Printf("Hello %s", "Programming with Go ☯\n")
}
