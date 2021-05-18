// package leap contains simple functions about leap years
package leap

// IsLeapYear checks if parameter year is a leap year
func IsLeapYear(year int) bool {
	evenlyDivisibleBy4 := evenlyDivisibleBy(year, 4)
	evenlyDivisibleBy100 := evenlyDivisibleBy(year, 100)
	evenlyDivisibleBy400 := evenlyDivisibleBy(year, 400)

	return (evenlyDivisibleBy4 && !evenlyDivisibleBy100) ||
		(evenlyDivisibleBy4 && evenlyDivisibleBy400)
}

func evenlyDivisibleBy(dividend, divisor int) bool {
	return dividend%divisor == 0
}
