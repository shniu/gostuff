package cn

//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 
//
// 两个字符串完全匹配才算匹配成功。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。 
// 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
// 
//
// 示例 3: 
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
// 
//
// 示例 4: 
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 
//
// 示例 5: 
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输入: false 
// Related Topics 贪心算法 字符串 动态规划 回溯算法

//leetcode submit region begin(Prohibit modification and deletion)

// 通配符匹配典型解法有两个，双指针法和动态规划

// 双指针
func isMatch1(s string, p string) bool {
	var sp, pp, match int
	var star = -1

	for sp < len(s) {
		// ? or s[sp] == p[pp]
		if pp < len(p) && (p[pp] == '?' || s[sp] == p[pp]) {
			sp++
			pp++
		} else if pp < len(p) && p[pp] == '*' { // *
			match = sp
			star = pp
			pp++
		} else if star != -1 { //
			pp = star + 1 // 回溯
			match++
			sp = match
		} else {
			return false
		}
	}

	for pp < len(p) && p[pp] == '*' {
		pp++
	}
	return pp == len(p)
}

// dp
func isMatch(s string, p string) bool {
	var m, n = len(s), len(p)

	// dp[i][j] 表示s[0..i]和p[0..j]是否匹配
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	// base case
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 如果i, j对应位置的字符相同或者j对应位置是?通配符
			// DP 方程：dp[i][j] = dp[i-1][j-1]
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// 如果j对应位置是*
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
