// Package triangle have functions to verify triange properties and do triangle calculations
package triangle

import (
	"math"
	"sort"
)

type Kind int

const (
	NaT = iota
	Equ = iota
	Iso = iota
	Sca = iota
)

// KindFromSides returns the kind of an triangle based on its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !isNaturalNumber(a) || !isNaturalNumber(b) || !isNaturalNumber(c) {
		return NaT
	}
	sides := sidesOrderedBySize(a, b, c)
	switch {
	case isEquilateral(sides):
		k = Equ
	case isScalene(sides):
		k = Sca
	case isIsosceles(sides):
		k = Iso
	default:
		k = NaT
	}

	return k
}

func sidesOrderedBySize(p ...float64) []float64 {
	sort.Float64s(p)
	return p
}

func isNaturalNumber(n float64) bool {
	return n > 0 && !math.IsInf(n, 0)
}

func isIsosceles(sides []float64) bool {
	return sides[0]+sides[1] >= sides[2]
}

func isEquilateral(sides []float64) bool {
	return sides[0] == sides[1] &&
		sides[1] == sides[2]
}

func isScalene(sides []float64) bool {
	return sides[0] != sides[1] &&
		sides[0] != sides[2] &&
		sides[1] != sides[2] &&
		sides[0]+sides[1] >= sides[2]
}
