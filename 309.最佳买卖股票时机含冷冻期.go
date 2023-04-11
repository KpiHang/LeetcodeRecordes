/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */

// @lc code=start
// 代录；动态规划；
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0] = -prices[0]
	// dp[0][1] = 0 // 默认值为0；
	// dp[0][2] = 0
	// dp[0][3] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][3]-prices[i], dp[i-1][1]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}
	return max(dp[len(prices)-1][1], max(dp[len(prices)-1][2], dp[len(prices)-1][3]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

