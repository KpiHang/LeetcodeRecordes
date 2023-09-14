/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

/*
灵神，相向双指针2
数组无序；
*/
func maxArea(height []int) int {
	max_area := 0
	left, right := 0, len(height)-1
	for left < right {
		cur_area := (right - left) * min(height[left], height[right])
		max_area = max(max_area, cur_area)
		// 矮边移动，去找有没有其他边，让结果更大；
		// 对于高边来说，是贪心的；
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max_area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

