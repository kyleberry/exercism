// Package isogram returns true if a string is an isogram and false otherwise

package isogram

import (
	"strings"
)

// IsIsogram removes spaces and hypens, converts the input string to a rune
// slice and iterates across the array, adding it to a map type bool to filter
// out duplicates. Then if the length of our map and the length of our slice are
// equal we didn't filter out any characters so it must be an isogram. If the
// two lengths are not equal we return false.
func IsIsogram(input string) bool {
	lower := strings.ToLower(strings.TrimSpace(
		strings.Replace(strings.Replace(input, "-", " ", -1),
			" ", "", -1)))
	letters := []rune(lower)
	lettersSeen := make(map[rune]bool)
	for i := 0; i < len(letters); i++ {
		lettersSeen[letters[i]] = true
	}
	if len(letters) == len(lettersSeen) {
		return true
	}
	return false
}
