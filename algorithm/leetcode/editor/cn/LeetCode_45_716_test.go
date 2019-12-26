package cn

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	arr := []int{1,2,1,1,1}
	min := jump(arr)
	assert.Equal(t, 3, min)
}
