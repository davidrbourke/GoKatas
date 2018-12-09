package stringkatas

import (
	"testing"
)

func TestLongestConsecShouldFind(t *testing.T) {
	input := []string{"zone", "abigail", "theta", "form", "libe", "zas"}
	expectedResult := "abigailtheta"
	length := 2

	actualResult := LongestConsec(input, length)

	if actualResult != expectedResult {
		t.Errorf("Did not find correct string")
	}
}

func TestLongestConsecSingleWord(t *testing.T) {
	input := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	expectedResult := "oocccffuucccjjjkkkjyyyeehh"
	length := 1

	actualResult := LongestConsec(input, length)

	if actualResult != expectedResult {
		t.Errorf("Did not find correct string, found" + actualResult)
	}
}
