/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
// self; DP完全背包；
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	// 求最小，所以其他位置初始化为int的最大值；
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
		// fmt.Println(dp)
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

