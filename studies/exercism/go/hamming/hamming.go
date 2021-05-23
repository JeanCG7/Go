// Package hamming contains functions to calculate hamming differences
package hamming

import "errors"

// Distance returns the hammingDistance given two sequences of same length
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) != len(runesB) {
		return 0, errors.New("can't calculate hamming distance between sequences of different length")
	}

	hammingDistance := 0
	for i := 0; i < len(runesA); i++ {
		if runesA[i] != runesB[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
