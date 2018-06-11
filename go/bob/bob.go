// Package bob parses input to determine Bob's responses.
package bob

import (
	"strings"
	"unicode"
)

// Hey parses input from Alice and gives Bob's response
func Hey(remark string) string {
	var (
		// Alice's remark
		question = strings.HasSuffix(strings.TrimSpace(remark), "?")
		shouting = isShouting(remark)
		nothing  = strings.TrimSpace(remark) == ""
		// Bob's responses
		sure     = "Sure."
		woah     = "Whoa, chill out!"
		calm     = "Calm down, I know what I'm doing!"
		fine     = "Fine. Be that way!"
		whatever = "Whatever."
	)

	switch {
	case question && shouting:
		return calm
	case question:
		return sure
	case shouting:
		return woah
	case nothing:
		return fine
	default:
		return whatever
	}
}

func isShouting(remark string) bool {
	var hasLetter bool
	for _, letter := range remark {
		if unicode.IsLetter(letter) {
			hasLetter = true
			if unicode.IsLower(letter) {
				return false
			}
		}
	}
	return hasLetter
}
