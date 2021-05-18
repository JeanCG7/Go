//package acronym contains functions to transform phrases to acronyms
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns an acronym given a string
func Abbreviate(s string) string {
	var acronym string
	for {
		if s == "" {
			break
		}

		firsLetter := rune(s[0])
		index := getFirstNonLetterIndex(s)

		if !unicode.IsLetter(firsLetter) {
			s = replaceNonLetterChar(s, index)
			continue
		}

		acronym = acronym + strings.ToUpper((string(firsLetter)))

		if index == -1 {
			break
		}

		s = s[index+1:]
	}
	return acronym
}

func getFirstNonLetterIndex(phrase string) int {
	for i, s := range phrase {
		if !unicode.IsLetter(s) && string(s) != `'` {
			return i
		}
	}
	return -1
}

func replaceNonLetterChar(s string, index int) string {
	if index == 0 {
		return s[index+1:]
	} else if index == len(s) {
		return s[:index-1]
	}
	return s[:index-1] + s[index+1:]
}
