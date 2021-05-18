// Package bob contains simple functions to have an conversation with bob
package bob

import (
	"regexp"
	"strings"
	"unicode"
)

// Hey function returns bob's answer given an string
func Hey(remark string) string {
	if isEmptyString(remark) {
		return "Fine. Be that way!"
	}

	allUpperCase := allUpper(remark) && !allNonLetter(remark)
	isAnQuestion := isAnQuestion(remark)

	if allUpperCase && isAnQuestion {
		return "Calm down, I know what I'm doing!"
	} else if allUpperCase {
		return "Whoa, chill out!"
	} else if isAnQuestion {
		return "Sure."
	}

	return "Whatever."
}
func isAnQuestion(word string) bool {
	word = strings.Trim(word, " ")
	wordLastChar := string(word[len(word)-1])

	return wordLastChar == "?"
}

func allUpper(word string) bool {
	for _, r := range word {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func allNonLetter(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func isEmptyString(word string) bool {
	var re = regexp.MustCompile(`\r\n\t|[\t\r\n\v\f\x{0085}\x{2028}\x{2029}]`)
	word = re.ReplaceAllString(word, "")
	return len(strings.Trim(word, " ")) == 0
}
