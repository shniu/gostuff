package cn

//翻转一棵二叉树。
//
// 示例： 
//
// 输入： 
//
//      4
//   /   \
//  2     7
// / \   / \
//1   3 6   9 
//
// 输出： 
//
//      4
//   /   \
//  7     2
// / \   / \
//9   6 3   1 
//
// 备注: 
//这个问题是受到 Max Howell 的 原问题 启发的 ： 
//
// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。 
// Related Topics 树 
// 👍 798 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		if curr == nil {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		curr.Left, curr.Right = curr.Right, curr.Left

		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}

		curr = curr.Left
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
