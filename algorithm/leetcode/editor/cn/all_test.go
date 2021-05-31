package cn

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	p := &TreeNode{-10, nil, nil}
	p.Left = &TreeNode{9, nil, nil}
	p.Right = &TreeNode{
		20,
		&TreeNode{15, nil, nil},
		&TreeNode{7, nil, nil},
	}

	sum := maxPathSum(p)
	fmt.Println(sum)
}
