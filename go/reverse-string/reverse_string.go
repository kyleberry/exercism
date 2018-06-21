// Package reverse reverses a string
package reverse

// String converts input to runes, then iterates across the slice backwards,
// creating a new slice, then converts the new slice back to a string.
func String(input string) string {
	var reversed []rune
	runes := []rune(input)
	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}
