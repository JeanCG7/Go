// Package isogram contains functons to check isogram existence
package isogram

import "unicode"

// func IsIsogram(word string) bool {
// 	m := make(map[rune]int)

// 	for _, v := range word {
// 		v = unicode.ToUpper(v)
// 		if m[v] == 1 && unicode.IsLetter(v) {
// 			return false
// 		}
// 		m[v] = 1
// 	}
// 	return true
// }

//IsIsogram checks if an string input is an isogram
func IsIsogram(word string) bool {
	letters := make([]rune, 0)

	for _, v := range word {
		if unicode.IsLetter(v) {
			v = unicode.ToUpper(v)
			if containsLetter(letters, v) {
				return false
			}
			letters = append(letters, v)
		}
	}

	return true
}

func containsLetter(word []rune, letter rune) bool {
	for _, v := range word {
		if v == letter {
			return true
		}
	}
	return false
}
