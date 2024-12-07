package parser

import (
	"fmt"
	"testing"
)

func TestSimpleThematicBreak(t *testing.T) {
	lines := []string{"***", "---", "___"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() != "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block should be a thematic break but is instead %v", block.GetNodeType())
		}
	}
}

func TestThematicBreakWithTooManySpacesBefore(t *testing.T) {
	lines := []string{"    ***", "     ---"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() == "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block shouldn't be a thematic break because it has too many spaces")
		}
	}
}

func TestThematicBreakWithMoreChars(t *testing.T) {
	lines := []string{"******************************************", "----------"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() != "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block should be a thematic break but is instead %v", block.GetNodeType())
		}
	}
}

func TestThematicBreakWithSpacesBetweenChars(t *testing.T) {
	lines := []string{"* * * * * *  ", "   *          *             *"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() != "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block should be a thematic break but is instead %v", block.GetNodeType())
		}
	}
}

func TestThematicBreakWithTabsBetweenChars(t *testing.T) {
	lines := []string{"*	*	*	* * *  ", "  -		- ---"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		fmt.Printf("%T\n", block)
		if block.GetNodeType() != "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block should be a thematic break but is instead %v", block.GetNodeType())
		}
	}
}

func TestThematicBreakWithOtherCharsOnLine(t *testing.T) {
	lines := []string{"* * * *a * *  ", "------a"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() == "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block shouldn't be a thematic break because it has a different char on the line")
		}
	}
}

func TestThematicBreakWithDifferentValidChars(t *testing.T) {
	lines := []string{"**---*"}
	root := Parse(lines)

	for _, block := range root.GetChildren() {
		if block.GetNodeType() == "THEMATIC_BREAK_BLOCK" {
			t.Errorf("The block shouldn't be a thematic break because it has a different char on the line")
		}
	}
}

func TestThematicBreakBreakingParagraph(t *testing.T) {
	lines := []string{"Paragraph 1", "***", "Paragraph 2"}
	root := Parse(lines)

	block := root.GetChildren()[1]

	if block.GetNodeType() != "THEMATIC_BREAK_BLOCK" {
		t.Errorf("The block should be a thematic block but is instead %v", block.GetNodeType())
	}
}
