/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
// Self: 完全背包；
func numSquares(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for j := 1; j <= n; j++ {
		dp[j] = math.MaxInt32
	}
	for j := 1; j <= n; j++ {
		for i := 0; i < n; i++ {
			if i*i <= j {
				dp[j] = min(dp[j], dp[j-i*i]+1)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

