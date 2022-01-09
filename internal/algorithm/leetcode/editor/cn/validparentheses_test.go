package cn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses_simple(t *testing.T) {
	s := "()"
	r := isValidParentheses(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_complex(t *testing.T) {
	s := "[[[{([])}]]]"
	r := isValidParentheses(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_empty(t *testing.T) {
	s := ""
	r := isValidParentheses(s)

	assert.Equal(t, true, r)
}

func TestValidParentheses_error(t *testing.T) {
	s := "{{{([[]])}]}"
	r := isValidParentheses(s)

	assert.Equal(t, false, r)
}
