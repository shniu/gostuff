package cn

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var val []int
	val = append(val, 123)
	val = append(val, 222)

	p(val)

	for _, v := range val {
		fmt.Println(v)
	}
}

func p(s []int) {
	s[0] = 100
}

func TestPathSum(t *testing.T) {
	p := &TreeNode{5, nil, nil}
	p.Left = &TreeNode{4, nil, nil}
	p.Right = &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}

	res := pathSum(p, 14)
	fmt.Println(res)
}

func TestFlatten(t *testing.T) {
	p := &TreeNode{5, nil, nil}
	p.Left = &TreeNode{4, nil, nil}
	p.Right = &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}

	flatten(p)

	fmt.Println(p)
}