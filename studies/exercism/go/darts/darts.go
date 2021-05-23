// Package darts contains functions to calculate points earned in a dart game
package darts

import (
	"math"
)

// Score returns the earned points from a toss in a dart game given an x and y axis
func Score(x, y float64) int {
	switch r := math.Sqrt(x*x + y*y); {
	case 1.0 >= r:
		return 10
	case 5.0 >= r:
		return 5
	case 10.0 >= r:
		return 1
	default:
		return 0
	}
}
