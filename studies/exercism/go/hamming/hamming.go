// Package hamming contains functions to calculate hamming differences
package hamming

import "errors"

// Distance returns the hammingDistance given two sequences of same length
func Distance(a, b string) (int, error) {
	var hammingDistance int
	if len(a) != len(b) {
		return 0, errors.New("can't calculate hamming distance between sequences of different length")
	}

	for i := range a {
		if a[i] != b[i] {
			hammingDistance = hammingDistance + 1
		}
	}

	return hammingDistance, nil
}
