// Package hamming calculates the Hamming Distance between two DNA strands
package hamming

import "errors"

// Distance calculates the Hamming Distance between two DNA strands and returns
// an error if the length of the strands is not equal.
func Distance(a, b string) (int, error) {
	var hammingDistance int

	if len(a) != len(b) {
		return 0, errors.New("DNA strands must be of equal length.")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
