package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterString(t *testing.T) {
	IterString_fori("abc")

	IterString_range("xyz")
}

func TestStrToInt(t *testing.T) {
	i := StrToInt("44")
	assert.Equal(t, 44, i)

	i2 := StrToInt("4a")
	assert.Equal(t, -1, i2)
}
