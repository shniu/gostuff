package cn

import "math"

//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
//一定经过根节点。 
//
// 路径和 是路径中各节点值的总和。 
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6 
//
// 示例 2： 
//
// 
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
// 
//
// 
//
// 提示： 
//
// 
// 树中节点数目范围是 [1, 3 * 104] 
// -1000 <= Node.val <= 1000 
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
func maxPathSum(root *TreeNode) int {
	var maxSumVal = math.MinInt32

	var nodeMaxGain func(root *TreeNode) int
	nodeMaxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		maxLeft := maxSum(nodeMaxGain(root.Left), 0);
		maxRight := maxSum(nodeMaxGain(root.Right), 0);

		newPathGain := maxLeft + maxRight + root.Val

		maxSumVal = maxSum(newPathGain, maxSumVal)

		return root.Val + maxSum(maxRight, maxLeft)
	}

	nodeMaxGain(root)
	return maxSumVal
}

func maxSum(x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)

// 下面的解法可以优化
//var maxSumVal = math.MinInt32
//func maxPathSum(root *TreeNode) int {
//	nodeMaxGain(root)
//	res := maxSumVal
//	maxSumVal = math.MinInt32
//	return res
//}
//
//func nodeMaxGain(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	maxLeft := maxSum(nodeMaxGain(root.Left), 0);
//	maxRight := maxSum(nodeMaxGain(root.Right), 0);
//
//	newPathGain := maxLeft + maxRight + root.Val
//
//	maxSumVal = maxSum(newPathGain, maxSumVal)
//
//	return root.Val + maxSum(maxRight, maxLeft)
//}
//
//func maxSum(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
