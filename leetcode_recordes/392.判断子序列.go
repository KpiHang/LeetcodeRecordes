/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
// 法2：动态规划
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n] == len(s)
}

// 法1：双指针法
func isSubsequence_1(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i += 1
		}
		j += 1
	}
	// 如果true，一定是遍历了s的所有字符
	return i == len(s)
}

// @lc code=end

