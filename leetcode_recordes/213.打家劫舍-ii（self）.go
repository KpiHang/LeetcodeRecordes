/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
// 代录-Self；动态规划；打家劫舍系列；
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp_L, dp_R := make([]int, len(nums)), make([]int, len(nums))
	dp_L[0], dp_L[1] = 0, nums[0]
	for i := 2; i < len(nums); i++ {
		dp_L[i] = max(dp_L[i-1], dp_L[i-2]+nums[i-1])
	}
	dp_R[0], dp_R[1] = 0, nums[1]
	for i := 2; i < len(nums); i++ {
		dp_R[i] = max(dp_R[i-1], dp_R[i-2]+nums[i])
	}
	return max(dp_L[len(nums)-1], dp_R[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

