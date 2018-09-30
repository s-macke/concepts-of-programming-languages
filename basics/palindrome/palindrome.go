// Copyright 2018 Johannes Weigend, Johannes  Siedersleben
// Licensed under the Apache License, Version 2.0

// Package palindrome implements multiple functions for palindromes.
package palindrome

import "github.com/jweigend/concepts-of-programming-languages/basics/types/strings"

// IsPalindrome implementation
// START OMIT
func IsPalindrome(word string) bool {
	for pos := 0; pos < len(word)/2; pos++ {
		if word[pos] != word[len(word)-pos-1] {
			return false
		}
	}
	return true
}

// END OMIT

// IsPalindrome2 is implemented by using runes.
func IsPalindrome2(word string) bool {
	var runes = []rune(word)
	for pos, ch := range runes {
		if ch != runes[len(runes)-pos-1] {
			return false
		}
	}
	return true
}

// IsPalindrome3 is implemented by reusing Reverse()
func IsPalindrome3(word string) bool {
	return strings.Reverse(word) == word
}
