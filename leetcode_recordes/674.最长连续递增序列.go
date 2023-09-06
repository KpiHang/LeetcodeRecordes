/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
// Self 动态规划；
func findLengthOfLCIS(nums []int) (ans int) {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			ans = max(ans, dp[i-1])
			dp[i] = 1
		} else {
			dp[i] = dp[i-1] + 1
		}
	}
	return max(ans, dp[len(nums)-1])
}

// 官方题解 双指针法；
func findLengthOfLCIS_1(nums []int) (ans int) {
	start := 0
	for i, v := range nums {
		if i > 0 && v <= nums[i-1] {
			start = i
		}
		ans = max(ans, i-start+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

