package cn

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	treeNodes := generateTrees(2)

	fmt.Println("Size: ", len(treeNodes))

	for node := range treeNodes {
		fmt.Println(node)
	}
}

func TestNumTrees(t *testing.T) {
	fmt.Println("Num 10, tree nums is ", numTrees(10))
}

func TestIsValidBST(t *testing.T) {
	// [5,4,6,null,null,3,7]
	root := &TreeNode{5, nil, nil}
	root.Left = &TreeNode{4, nil, nil}
	root.Right = &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}

	fmt.Println("Expected false, return is ", isValidBST(root))
}

func TestSameTree(t *testing.T) {
	p := &TreeNode{5, nil, nil}
	p.Left = &TreeNode{4, nil, nil}
	p.Right = &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}

	q := &TreeNode{5, nil, nil}
	q.Left = &TreeNode{4, nil, nil}
	q.Right = &TreeNode{
		6,
		&TreeNode{3, nil, nil},
		&TreeNode{7, nil, nil},
	}

	b := isSameTree(p, q)
	fmt.Println("Result of same tree: ", b)
}