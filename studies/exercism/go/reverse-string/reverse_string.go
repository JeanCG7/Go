// Package reverse contains functions to create reverse strings
package reverse

// Reverse returns the reverse of a given string
func Reverse(word string) string {
	runes := []rune(word)

	reverseWord := ""
	for i := len(runes) - 1; i >= 0; i-- {
		reverseWord += string(runes[i])
	}

	return reverseWord
}
