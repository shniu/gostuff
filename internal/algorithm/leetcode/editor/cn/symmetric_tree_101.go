package cn
//给定一个二叉树，检查它是否是镜像对称的。 
//
// 
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。 
//
//     1
//   / \
//  2   2
// / \ / \
//3  4 4  3
// 
//
// 
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的: 
//
//     1
//   / \
//  2   2
//   \   \
//   3    3
// 
//
// 
//
// 进阶： 
//
// 你可以运用递归和迭代两种方法解决这个问题吗？ 
// Related Topics 树 深度优先搜索 广度优先搜索 
// 👍 1298 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)

	for len(queue) > 0 {
		node1 := queue[0]
		queue = queue[1:]
		node2 := queue[0]
		queue = queue[1:]

		if node1 == nil && node2 == nil {
			continue
		}

		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}

		queue = append(queue, node1.Left)
		queue = append(queue, node2.Right)
		queue = append(queue, node1.Right)
		queue = append(queue, node2.Left)
	}

	return true
}
//leetcode submit region end(Prohibit modification and deletion)

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSame(root.Left, root.Right)
}

func isSame(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if (p == nil && q != nil) || (p != nil && q == nil) || p.Val != q.Val {
		return false
	}

	return p.Val == q.Val && isSame(p.Left, q.Right) && isSame(p.Right, q.Left)
}
