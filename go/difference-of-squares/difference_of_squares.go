// Package diffsquares determines the difference between the sum of squares and
// the square of sums for the first `n` natural numbers.
package diffsquares

// SquareOfSums finds the square of the sum of the first `n` natural numbers
func SquareOfSums(input int) int {
	var sqOfSums float64
	var sum int
	for i := 0; i <= input; i++ {
		sum += i
		sqOfSums = float64(sum * sum)
	}
	return int(sqOfSums)
}

// SumOfSquares finds the sum of the squares of the first `n` naural numbers.
func SumOfSquares(input int) int {
	var sumOfSq float64
	for i := 0; i <= input; i++ {
		sumOfSq += float64(i * i)
	}
	return int(sumOfSq)
}

// Difference finds the difference between the square of sums and sum of squares
// of the first `n` nautral numbers.
func Difference(input int) int {
	return int(SquareOfSums(input) - SumOfSquares(input))
}
