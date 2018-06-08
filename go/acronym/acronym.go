// Package acronym returns an acronym from input
package acronym

import (
	"regexp"
	"strings"
	"unicode"
)

// removeLower uses srings.Map to iterate over a string, removing lower case using unicode.IsLower
func removeLower(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return -1
		}
		return r
	}, s)
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	// Make everything lowercase first
	lower := strings.ToLower(s)
	// Convert the first letter of each word to Upper
	titleize := strings.Title(lower)
	// Remove non alpha numeric characters
	nonAlphaStr := regexp.MustCompile(`[^[:alnum:]\s]`).ReplaceAllString(titleize, "")
	// Remove whitespace
	noWhitespace := strings.Replace(nonAlphaStr, " ", "", -1)
	// Finally return our acronym by removing lowercase letters in removeLower
	return removeLower(noWhitespace)
}
