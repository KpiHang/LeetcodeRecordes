/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
// 代录；DP；两种状态0：持有；1：不持有
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

