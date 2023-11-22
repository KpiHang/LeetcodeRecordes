/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start

// 滑动窗口，寻找最短；
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum, shortest := 0, len(nums)+1
	for right < len(nums) {
		sum += nums[right]
		right++

		for sum >= target {
			if right-left < shortest {
				shortest = right - left
			}
			sum -= nums[left] // 可能会导致sum < target，刚好是滑动窗口下一个起始位置；
			left++
		}
	}

	if shortest == len(nums)+1 {
		return 0
	}
	return shortest
}

// @lc code=end

