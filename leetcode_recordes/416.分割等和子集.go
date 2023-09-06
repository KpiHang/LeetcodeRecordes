/*
 * @lc app=leetcode.cn id=416 lang=golang
 *
 * [416] 分割等和子集
 */

// @lc code=start
// 代录；0-1背包问题的一维数组解法；
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果sum为奇数，则不能平分为两个子集；
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return dp[target] == target
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end

