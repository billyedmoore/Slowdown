package utils

import "testing"

func TestRemoveWhitespaceSpaces(t *testing.T) {
	s := "this is a string"
	s = RemoveWhitespace(s)

	expected := "thisisastring"
	if s != expected {
		t.Fatalf("Expected %v but got %v instead", expected, s)
	}
}

func TestRemoveWhitespaceTabs(t *testing.T) {
	s := "this	is	a	string"
	s = RemoveWhitespace(s)

	expected := "thisisastring"
	if s != expected {
		t.Fatalf("Expected %v but got %v instead", expected, s)
	}

}
