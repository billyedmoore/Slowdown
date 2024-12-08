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

func TestDoesStartWithThreeOrLessSpaces(t *testing.T) {
	s := "   this is a string"
	does := DoesLineStartWithThreeOrLessSpaces(s)

	if does != true {
		t.Fatalf("Expected \"%v\" to be true got false instead", s)
	}
}

func TestDoesStartWithThreeOrLessSpacesTooManySpaces(t *testing.T) {
	s := "    this is a string"
	does := DoesLineStartWithThreeOrLessSpaces(s)

	if does != false {
		t.Fatalf("Expected \"%v\" to be false got true instead", s)
	}
}

func TestDoesStartWithThreeOrLessSpacesNoSpaces(t *testing.T) {
	s := "this is a string"
	does := DoesLineStartWithThreeOrLessSpaces(s)

	if does != true {
		t.Fatalf("Expected \"%v\" to be true got false instead", s)
	}
}
