package stringkatas

import "testing"

func TestHasUniqueChar_ShouldBeTrue(t *testing.T) {
	testStr := "abcdef"

	hasUniqueChar := HasUniqueChar(testStr)

	if !hasUniqueChar {
		t.Errorf("HasUniqueChar was false but should be true")
	}
}

func TestHasUniqueChar_ShouldBeFalse(t *testing.T) {
	testStr := "abcdefb"

	hasUniqueChar := HasUniqueChar(testStr)

	if hasUniqueChar {
		t.Errorf("HasUniqueChar was true but should be false")
	}
}
