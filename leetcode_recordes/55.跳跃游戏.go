/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[start]-(end-start) <= nums[end] {
			start = end
			if nums[start] == 0 && start != len(nums)-1 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

