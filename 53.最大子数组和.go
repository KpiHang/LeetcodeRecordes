/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
// 贪心算法
// https://leetcode.cn/problems/maximum-subarray/solutions/1536926/by-nehzil-rmeh/?topicTags=greedy
func maxSubArray(nums []int) int {
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

