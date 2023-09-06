/*
 * @lc app=leetcode.cn id=714 lang=golang
 *
 * [714] 买卖股票的最佳时机含手续费
 */

// @lc code=start
// Self; 动态规划；注意与122，123的区别；
// 0买入，1卖出；
func maxProfit(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

