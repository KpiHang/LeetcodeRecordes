/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
// Self；动态规划DP；
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1) // dp[i]: 到达第i层所需要的最小花费；
	dp[0], dp[1] = 0, 0            // 初始化；
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2]) // 递推公式；
	}
	return dp[len(cost)]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end
