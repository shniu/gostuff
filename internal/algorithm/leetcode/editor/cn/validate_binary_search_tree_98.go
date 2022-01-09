package cn

//给定一个二叉树，判断其是否是一个有效的二叉搜索树。
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索 递归 
// 👍 963 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST1(root *TreeNode) bool {
	var min = -1 << 63
	var max = 1<<63 - 1
	return isValidBSTHelper(root, min, max)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
}

func isValidBST(root *TreeNode) bool {
	var min = -1 << 63
	return isValid(root, &min)
}

func isValid(node *TreeNode, previous *int) bool {
	if node == nil {
		return true
	}

	leftValid := isValid(node.Left, previous)
	if !leftValid {
		return false
	}

	if *previous >= node.Val {
		return false
	} else {
		*previous = node.Val
	}

	rightValid := isValid(node.Right, previous)
	if !rightValid {
		return false
	}

	return leftValid && rightValid
}

//leetcode submit region end(Prohibit modification and deletion)

///
var orders []int

func isValidBST3(root *TreeNode) bool {
	orders = make([]int, 0)
	b := isValid3(root)

	//fmt.Println("=====")
	//for entry := range orders {
	//	fmt.Println(entry, orders[entry])
	//}

	return b
}

func isValid3(node *TreeNode) bool {
	if node == nil {
		return true
	}

	leftValid := isValid3(node.Left)
	if !leftValid {
		return false
	}

	if len(orders) >= 1 && orders[len(orders) - 1] >= node.Val {
		return false
	} else {
		orders = append(orders, node.Val)
	}
	// fmt.Println(node)
	rightValid := isValid3(node.Right)
	if !rightValid {
		return false
	}

	return leftValid && rightValid
}
