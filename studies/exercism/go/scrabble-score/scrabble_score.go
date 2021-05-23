// Package scrabble contains functions to return values given input strings
package scrabble

import "unicode"

var scores = map[rune]int{}

func init() {
	initScores := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}

	for r, v := range initScores {
		for _, letter := range r {
			scores[letter] = v
		}
	}
}

// Score returns a score given an string
func Score(word string) int {
	if word == "" {
		return 0
	}

	result := 0
	for _, letter := range word {
		result += scores[unicode.ToUpper(letter)]
	}

	return result
}
