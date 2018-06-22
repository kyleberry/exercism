// Package luhn validates a number using the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid validates a number using the Luhn algorithm
func Valid(input string) bool {
	var sum int

	numbers := strings.Replace(input, " ", "", -1)

	if len(numbers) <= 1 {
		return false
	}

	for _, char := range numbers {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	for i := 0; i < len(numbers); i++ {
		digit, _ := strconv.Atoi(numbers[i : i+1])

		if (len(numbers)-i)%2 == 0 {
			double := digit * 2
			sum += double/10 + double%10
		} else {
			sum += digit
		}
	}

	return sum >= 0 && sum%10 == 0
}
