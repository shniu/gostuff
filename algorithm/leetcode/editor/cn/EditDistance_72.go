package cn

//给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。
//
// 你可以对一个单词进行如下三种操作： 
//
// 
// 插入一个字符 
// 删除一个字符 
// 替换一个字符 
// 
//
// 示例 1: 
//
// 输入: word1 = "horse", word2 = "ros"
//输出: 3
//解释: 
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')
// 
//
// 示例 2: 
//
// 输入: word1 = "intention", word2 = "execution"
//输出: 5
//解释: 
//intention -> inention (删除 't')
//inention -> enention (将 'i' 替换为 'e')
//enention -> exention (将 'n' 替换为 'x')
//exention -> exection (将 'n' 替换为 'c')
//exection -> execution (插入 'u')
// 
// Related Topics 字符串 动态规划

//leetcode submit region begin(Prohibit modification and deletion)

// DP 解法
func minDistance1(word1 string, word2 string) int {
	var m = len(word1)
	var n = len(word2)

	//var dp [][]int
	//for i := 0; i <= m; i++ {
	//	tmp := make([]int, n+1)
	//	dp = append(dp, tmp)
	//}

	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// base case
	dp[0][0] = 0
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

// 递归解法: 记忆化递归
func minDistance2(word1 string, word2 string) int {
	var m = len(word1)
	var n = len(word2)
	memo := make([][]int, m)
	for i := 0; i < m+1; i++ {
		memo[i] = make([]int, n)
	}

	return minDistanceHelper(word1, word2, memo, m-1, n-1)
}

func minDistanceHelper(word1, word2 string, memo [][]int, i, j int) int {
	// terminator
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	// process current logic
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	if word1[i] == word2[j] {
		// drill down
		memo[i][j] = minDistanceHelper(word1, word2, memo, i-1, j-1)
	} else {
		memo[i][j] = min(min(minDistanceHelper(word1, word2, memo, i-1, j-1),
			minDistanceHelper(word1, word2, memo, i, j-1)),
			minDistanceHelper(word1, word2, memo, i-1, j)) + 1
	}
	return memo[i][j]
}

//leetcode submit region end(Prohibit modification and deletion)
