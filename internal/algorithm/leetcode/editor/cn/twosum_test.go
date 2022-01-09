package cn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var nums = []int{2, 4, 10, 22}
	res := twoSum(nums, 14)
	fmt.Println(len(res))
	fmt.Println(res)

	assert.Equal(t, 2, len(res))
	assert.Equal(t, 1, res[0])
	assert.Equal(t, 2, res[1])
}
