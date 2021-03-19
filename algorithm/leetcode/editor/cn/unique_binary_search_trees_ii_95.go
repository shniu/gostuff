package cn
//给定一个整数 n，生成所有由 1 ... n 为节点所组成的 二叉搜索树 。 
//
// 
//
// 示例： 
//
// 输入：3
//输出：
//[
//  [1,null,3,2],
//  [3,2,null,1],
//  [3,1,null,null,2],
//  [2,1,3],
//  [1,null,2,null,3]
//]
//解释：
//以上的输出对应以下 5 种不同结构的二叉搜索树：
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
// 
//
// 
//
// 提示： 
//
// 
// 0 <= n <= 8 
// 
// Related Topics 树 动态规划 
// 👍 815 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return construct(1, n)
}

func construct(start, end int) []*TreeNode {

	// terminator
	if start > end {
		return []*TreeNode{}
	}

	var treeNodes []*TreeNode

	for i := start; i <= end; i++ {
		leftNodes := construct(start, i - 1)
		rightNodes := construct(i + 1, end)

		for _, left := range leftNodes {
			for _, right := range rightNodes {
				curr := &TreeNode{i, nil, nil}
				curr.Left = left
				curr.Right = right

				treeNodes = append(treeNodes, curr)
			}
		}
	}

	return treeNodes
}
//leetcode submit region end(Prohibit modification and deletion)
