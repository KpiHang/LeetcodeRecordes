/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
// 代录；DP：打家劫舍系列；
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	result1 := robRange(nums[0 : len(nums)-1])
	result2 := robRange(nums[1:])
	return max(result1, result2)
}

func robRange(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

