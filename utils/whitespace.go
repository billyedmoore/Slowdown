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
