// Package bob should have a package comment that summarizes what it's about.
package bob

import "regexp"

// Parses input from Alice and gives Bob's response
func Hey(remark string) string {
	var (
		// Alice's remark regexp
		question         = regexp.MustCompile(`^[\w\W\s]+\?\s*$`)
		shouting         = regexp.MustCompile(`^[A-Z0-9\W\s]+$`)
		shoutingQuestion = regexp.MustCompile(`^[A-Z0-9\W\s]+\?\s*$`)
		forceful         = regexp.MustCompile(`[A-Z]`)
		nothing          = regexp.MustCompile(`^\s*$`)
		// Bob's responses
		sure     = "Sure."
		woah     = "Whoa, chill out!"
		calm     = "Calm down, I know what I'm doing!"
		fine     = "Fine. Be that way!"
		whatever = "Whatever."
	)

	switch {
	case shoutingQuestion.MatchString(remark) && forceful.MatchString(remark):
		return calm
	case question.MatchString(remark):
		return sure
	case shouting.MatchString(remark) && forceful.MatchString(remark):
		return woah
	case nothing.MatchString(remark):
		return fine
	default:
		return whatever
	}

}
