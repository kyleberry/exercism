// Package reverse reverses a string
package reverse

// String converts input to runes, then iterates across the slice backwards,
// creating a new slice, then converts the new slice back to a string.
func String(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
