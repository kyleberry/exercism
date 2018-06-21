// Package scrabble

package scrabble

import (
	"strings"
)

// Score does things
func Score(input string) int {
	var (
		letterScores      = make(map[string]int, 26)
		onePointLetters   = [10]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
		twoPointLetters   = [2]string{"D", "G"}
		threePointLetters = [4]string{"B", "C", "M", "P"}
		fourPointLetters  = [5]string{"F", "H", "V", "W", "Y"}
		fivePointLetters  = "K"
		eightPointLetters = [2]string{"J", "X"}
		tenPointLetters   = [2]string{"Q", "Z"}
		score             int
	)
	// Generate a map of scores per letter
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
	letterScores[fivePointLetters] = 5
	for i := 0; i < len(eightPointLetters); i++ {
		letterScores[eightPointLetters[i]] = 8
	}
	for i := 0; i < len(tenPointLetters); i++ {
		letterScores[tenPointLetters[i]] = 10
	}

	// Calculate score of input sting
	for _, rune := range input {
		letter := string(rune)
		score += letterScores[strings.ToUpper(letter)]
	}
	return score
}
