package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEmptyString_returnEmpty(t *testing.T) {
	res := decodeString("")
	assert.Equal(t, "", res)
}

func TestOneDigitString(t *testing.T) {
	res := decodeString("10[a]")
	assert.Equal(t, "aaaaaaaaaa", res)

	res2 := decodeString("2[a]3[bb]4[c]")
	assert.Equal(t, "aabbbbbbcccc", res2)

	res3 := decodeString("2[a]3[bb2[v2[x]]]4[c]")
	assert.Equal(t, "aabbvxxvxxbbvxxvxxbbvxxvxxcccc", res3)
}
