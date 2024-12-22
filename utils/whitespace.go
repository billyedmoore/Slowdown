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
	return HowManySpacesDoesLineStartWith(s) <= 3
}

func HowManySpacesDoesLineStartWith(s string) int {
	spaces := 0

	for _, c := range s {
		if c == ' ' {
			spaces += 1
		} else {
			break
		}
	}
	return spaces
}
