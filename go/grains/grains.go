// Package grains calculates how much grain out friend is going to end up with
// after saving the king's son.
package grains

import "errors"

// import "math"

// ! Slow readable way:
// BenchmarkSquare-8        5000000               364 ns/op
// BenchmarkTotal-8          500000              2694 ns/op

// Square calculates the number of grains of rice on a single chess board using
// ar^(n-1) where `a` is the starting position, `r` is the common factor between
// the numbers, and `n` is the position we're calculating.
// func Square(input int) (uint64, error) {
// 	if input > 64 || input < 1 {
// 		return 0, errors.New("There are only 64 squares on a chess board.")
// 	}
// 	return uint64(math.Pow(2, (float64(input - 1)))), nil
// }

// Total calculates the total number of grains of rice as defined in Square()
// func Total() uint64 {
// 	var totalGrainsOfRice uint64
// 	for i := 1; i <= 64; i++ {
// 		square, _ := Square(i)
// 		totalGrainsOfRice += square
// 	}
// 	return totalGrainsOfRice
// }

// ! Fast fancy bitwise magicks

// BenchmarkSquare-8       30000000                 43.4 ns/op
// BenchmarkTotal-8        2000000000               0.31 ns/op

// Square uses bit shift to avoid importing "math", the bitwise operator "<<"
// basically means multiply by 2 x number of times, so "n << x" would be
// "multiply n times 2 x times", in this exercism our common factor is 2, so we
// can use this as a shortcut.
func Square(input int) (uint64, error) {
	if input > 64 || input < 1 {
		return 0, errors.New("There are only 64 squares on a chess board.")
	}
	return 1 << uint64(input-1), nil
}

// Total is a function using magic bitwise operators, same as explained above.
// In this case we're saying "1 times 2 63 times" which is the same thing we did
// above using a for loop and some simple math.
func Total() uint64 {
	return 1<<64 - 1
}
