/*
 * @lc app=leetcode.cn id=518 lang=golang
 *
 * [518] 零钱兑换 II
 */

// @lc code=start
// Self; 完全背包问题；动态规划
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// 初始化为0
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coin]
		}
	}
	return dp[amount]
}

// @lc code=end

