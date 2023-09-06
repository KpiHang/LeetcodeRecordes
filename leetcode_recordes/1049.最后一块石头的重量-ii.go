/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
// 0-1背包，一维数组法；
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2

	dp := make([]int, target+1)
	dp[0] = 0
	for _, stone := range stones {
		for j := target; j >= stone; j-- { // 若 j < stones[j], 则物品i放不进去，也就没意义；
			dp[j] = max(dp[j], dp[j-stone]+stone)
		}
	}
	return sum - 2*dp[target]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

