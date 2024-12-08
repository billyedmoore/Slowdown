package utils

import "unicode"

func RemoveWhitespace(s string) string {
	returnVal := []rune{}

	for _, char := range []rune(s) {
		if !unicode.IsSpace(char) {
			returnVal = append(returnVal, char)
		}
	}

	return string(returnVal)
}

func DoesLineStartWithThreeOrLessSpaces(s string) bool {
	// space as defined by the commonmark spec a single U+0020 char
	allowedSpaces := 3

	for _, c := range s {
		if c == ' ' {
			if allowedSpaces > 0 {
				allowedSpaces -= 1
			} else {
				return false
			}
		} else {
			break
		}
	}
	return true
}
