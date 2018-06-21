// Package raindrops converts an int to "raindrop-speak"
package raindrops

import (
	"strconv"
	"strings"
)

// Convert takes on int as an input and converts it to "raindrop-speak"
// depending on the factor of the input, and outputs the result as a string.
func Convert(input int) string {
	output := " "

	if input%3 == 0 {
		output += "Pling"
	}
	if input%5 == 0 {
		output += "Plang"
	}
	if input%7 == 0 {
		output += "Plong"
	}
	if output == " " {
		output = strconv.Itoa(input)
	}

	return strings.TrimSpace(output)
}
