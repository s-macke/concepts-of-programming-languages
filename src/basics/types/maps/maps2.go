package main

import (
	"fmt"
	"strings"
)

// START OMIT
func countWords(input string) map[string]int {

	// Split the string into words using whitespace as a delimiter
	words := strings.Fields(input)

	// Initialize an empty map to store word counts
	wordCounts := make(map[string]int)

	// Iterate over the words and update the count in the map
	for _, word := range words {
		word = strings.ToLower(word) // Convert word to lowercase to ensure case-insensitivity
		wordCounts[word]++           // Increment the count for this word
	}

	return wordCounts
}

// END OMIT

func main() {
	text := "If it is to be it is up to me to delegate"
	wordCounts := countWords(text)

	fmt.Println("Word counts:")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
