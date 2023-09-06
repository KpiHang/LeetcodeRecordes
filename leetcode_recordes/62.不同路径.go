/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
// Self;DP;
func uniquePaths(m int, n int) int {
	// 定义dp数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 遍历；
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

