package stringkatas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInPlaceReverseWords_TwoWordsSameLength_Reversed(t *testing.T) {
	result := InPlaceReverseWords([]rune("test slow"))
	assert.Equal(t, "slow test", string(result))
}

func TestInPlaceReverseWords_FiveWordsDifferentLength_Reversed(t *testing.T) {
	result := InPlaceReverseWords([]rune("ok landed has eagle the"))
	assert.Equal(t, "the eagle has landed ok", string(result))
}

func TestInPlaceReverseWords_SingleWord_Unchanged(t *testing.T) {
	result := InPlaceReverseWords([]rune("test"))
	assert.Equal(t, "test", string(result))
}
