package diffsquares

// SumOfSquares given an int returns the sum of the squares from 1 to n
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSum given an int returns the square of the sum of 1 to n
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// Difference returns the absolute value from the difference between SumOfSquares and SquareofSum
func Difference(n int) int {
	a, b := SumOfSquares(n), SquareOfSum(n)
	if a < b {
		return b - a
	}

	return a - b
}
