// Package acronym returns an acronym from input
package acronym

import (
	"strings"
)

// Abbreviate removes separators, converts all letters to Upper, the pulls out
// the first letter of each word by splitting the string on whitespace using
// strings.Fields and range in a for loop.
func Abbreviate(s string) string {

	var abbreviation string

	noSeparator := strings.ToUpper(strings.Replace(s, "-", " ", -1))

	for _, words := range strings.Fields(noSeparator) {
		abbreviation += words[:1]
	}

	return abbreviation
}
