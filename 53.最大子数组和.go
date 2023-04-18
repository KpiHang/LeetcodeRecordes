/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
// 代录：DP法
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	// 初始化最大的和
	mx := nums[0]
	for i := 1; i < len(nums); i++ {
		// 会面临2种情况，一个是带前面的和，一个是不带前面的和
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		mx = max(mx, dp[i])
		// fmt.Println(dp)
	}
	return mx
}

// 贪心算法
// https://leetcode.cn/problems/maximum-subarray/solutions/1536926/by-nehzil-rmeh/?topicTags=greedy
func maxSubArray_贪心(nums []int) int {
	res := -int(math.Pow(10, 4)) - 1
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		res = max(res, tmp)
		if tmp < 0 {
			tmp = 0
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end

