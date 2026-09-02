package practice

import (
	"fmt"
	"strings"
)

/*
Exercise 018:

Validate passwords against these rules:
  1. At least one [a-z]
  2. At least one [A-Z]
  3. At least one [0-9]
  4. At least one of [$#@]
  5. Length between 6 and 12 (inclusive)
  6. No whitespace

Return the valid passwords from the comma-separated input, in the original
order, joined by commas.

Example:
  Ex018("ABd1234@1,a F1#,2w3E*,2We3345") -> "ABd1234@1"

Tip: run `go test ./...` from this folder.
*/

// Ex018 should return only the valid passwords, comma-separated.
func Ex018(input string) string {
	passwordSplits := strings.Split(input, ",")
	var output []string

	for _, val := range passwordSplits {
		//Length check
		if !LengthCheck(val, 6, 12) {
			continue
		}

		//'a'-'z' check
		if !CharacterRangeCheck(val, 'a', 'z') {
			continue
		}

		//'A' - 'Z' check
		if !CharacterRangeCheck(val, 'A', 'Z') {
			continue
		}

		//0-9 check
		if !CharacterRangeCheck(val, '0', '9') {
			continue
		}

		//$, #, @
		if !CharacterCheck(val, '$') && !CharacterCheck(val, '#') && !CharacterCheck(val, '@') {
			continue
		}

		//'Whitespace'
		if CharacterCheck(val, ' ') {
			continue
		}

		fmt.Println(val)
		output = append(output, val)
	}

	return strings.Join(output, ",")
}

func LengthCheck(input string, min, max int) bool {
	stringLen := len(input)
	if (stringLen >= 6) && (stringLen <= 12) {
		return true
	}
	return false
}

func CharacterRangeCheck(input string, min, max rune) bool {
	for _, sym := range input {
		if (sym >= min) && (sym <= max) {
			return true
		}
	}

	return false
}

func CharacterCheck(input string, char rune) bool {
	for _, sym := range input {
		if sym == char {
			return true
		}
	}

	return false
}
