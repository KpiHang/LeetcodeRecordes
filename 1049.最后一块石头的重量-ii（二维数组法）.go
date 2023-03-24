/*
 * @lc app=leetcode.cn id=1049 lang=golang
 *
 * [1049] 最后一块石头的重量 II
 */

// @lc code=start
// Self；0-1背包问题的二维数组解法；
func lastStoneWeightII(stones []int) int {
	// 找target；
	sum := 0
	for _, stone := range stones {
		sum += stone
	}
	target := sum / 2

	// 二维dp数组；
	dp := make([][]int, len(stones))
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]int, target+1)
	}

	// 初始化；
	for i := 0; i < len(stones); i++ {
		dp[i][0] = 0
	}
	for j := target; j > 0; j-- {
		if j < stones[0] {
			break
		}
		dp[0][j] = stones[0]
	}

	// 遍历；
	for i := 1; i < len(stones); i++ {
		for j := 1; j <= target; j++ {
			if j < stones[i] { // 背包的总容量，不能比当前的石头重量还小吧；
				dp[i][j] = dp[i-1][j] // 那只能考虑不放物品i的情况了；
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i])
			}
		}
	}
	res := sum - 2*dp[len(stones)-1][target]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

