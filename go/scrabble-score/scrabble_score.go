// Package scrabble

package scrabble

import (
	"strings"
)

// Score does things
func Score(input string) int {
	var (
		letterScores      = make(map[rune]int, 26)
		onePointLetters   = []rune("AEIOULNRST")
		twoPointLetters   = []rune("DG")
		threePointLetters = []rune("BCMP")
		fourPointLetters  = []rune("FHVWY")
		fivePointLetters  = []rune("K")
		eightPointLetters = []rune("JX")
		tenPointLetters   = []rune("QZ")
		score             int
	)
	// Generate a map of scores per rune
	for i := 0; i < len(onePointLetters); i++ {
		letterScores[onePointLetters[i]] = 1
	}
	for i := 0; i < len(twoPointLetters); i++ {
		letterScores[twoPointLetters[i]] = 2
	}
	for i := 0; i < len(threePointLetters); i++ {
		letterScores[threePointLetters[i]] = 3
	}
	for i := 0; i < len(fourPointLetters); i++ {
		letterScores[fourPointLetters[i]] = 4
	}
	letterScores[fivePointLetters[0]] = 5
	for i := 0; i < len(eightPointLetters); i++ {
		letterScores[eightPointLetters[i]] = 8
	}
	for i := 0; i < len(tenPointLetters); i++ {
		letterScores[tenPointLetters[i]] = 10
	}

	// Calculate score of input sting
	for _, rune := range strings.ToUpper(input) {
		score += letterScores[rune]
	}
	return score
}
