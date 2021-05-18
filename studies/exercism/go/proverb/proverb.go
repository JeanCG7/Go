// Package proverb has simple functions to generate proverbs based on different inputs
package proverb

import "fmt"

// Proverb returns an array of proverbs given a array of rhymes
func Proverb(rhyme []string) []string {
	var proverbs []string
	for i := range rhyme {
		if i == len(rhyme)-1 {
			proverbs = append(proverbs, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		} else {
			proverbs = append(proverbs, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	return proverbs
}
