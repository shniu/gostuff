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
// 使用动态规划求解
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	var treeNodes = make([][]*TreeNode, n + 1)
	treeNodes[0] = []*TreeNode{nil}
	treeNodes[1] = []*TreeNode{
		{1, nil, nil},
	}

	// 求 1...n 的所有组成的二叉搜索树，可以在所有 1...n-1 的基础上增加 n，并且加入以 n 为根节点的二叉搜索树
	// 需要利用从低向上的思想，从 1 开始，一直计算到 n，并且保留计算过程中的中间结果
	// 以 i 为根的二叉搜索树，1...i-1 在左，i+1...n 在右，继续分解 1...i-1 这个序列所有的二叉搜索树集合与i+1...n这个序列所有的二叉搜索树
	// 集合进行完全组合，就得到了以 i 为根节点的所有二叉搜索树的集合
	// 求 n 个数的序列时，需要把从 1 到 n 的所有二叉搜索树的结合都装入同一个集合中

	// 求序列 1...i 为根的二叉搜索树的集合
	for i := 2; i <= n; i++ {
		// 存储i的所有可能的二叉搜索树的集合
		treeNodes[i] = []*TreeNode{}

		// 分别以 1...i 为根节点，求出每种的二叉树集合
		for j := 1; j <= i; j++ {
			// 以 j 为根节点时，左节点可能的二叉搜索树集合
			leftNodes := treeNodes[j-1];
			// 以 j 为根节点时，右节点可能的二叉搜索树集合
			rightNodes := treeNodes[i-j]

			// 把左右集合做自由组合
			for k, _ := range leftNodes {
				for v, _ := range rightNodes {
					// 以 j 为根节点
					newNode := &TreeNode{j, nil, nil}
					// j 的左子树
					newNode.Left = leftNodes[k]
					// j 的右子树，由于结构完全是一样的，右子树每个节点的值要大一些，都加个 offset，然后把所有子树克隆一份
					newNode.Right = clone(rightNodes[v], j)

					// 追加到i的二叉搜索树的集合中
					treeNodes[i] = append(treeNodes[i], newNode)
				}
			}
		}
	}

	return treeNodes[n]
}

func clone(node *TreeNode, offset int) *TreeNode {
	if node == nil {
		return nil
	}

	cloneNode := &TreeNode{node.Val + offset, nil, nil}
	cloneNode.Left = clone(node.Left, offset)
	cloneNode.Right = clone(node.Right, offset)

	return cloneNode
}

func generateTrees2(n int) []*TreeNode {
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
