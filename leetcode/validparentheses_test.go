package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidParentheses_simple(t *testing.T) {
	s := "()"
	r := isValid(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_complex(t *testing.T) {
	s := "[[[{([])}]]]"
	r := isValid(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_empty(t *testing.T) {
	s := ""
	r := isValid(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_error(t *testing.T) {
	s := "{{{([[]])}]}"
	r := isValid(s)

	assert.Equal(t, false, r)
}
