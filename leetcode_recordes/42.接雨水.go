/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
/*
一个桶能装多少水，取决于左边木板最大高度和右边木板最大高度。
法 1：相向双指针 2
	一个桶能装多少水，取决于左边木板最大高度和右边木板最大高度。
	跨越间隔，不管多长，只和最高的相关；
	两边最高的夹起来就是最大能装多少；
*/
func trap(height []int) (ans int) {
	n := len(height)
	left, right := 0, n-1
	left_max, right_max := 0, 0
	for left < right {
		left_max = max(left_max, height[left])
		right_max = max(right_max, height[right])

		if left_max < right_max { // 短板效应；
			ans += left_max - height[left]
			left++
		} else { // 右侧是短板，右侧最高也就短板的容量了，所以右侧指针移动
			ans += right_max - height[right]
			right--
		}
	}
	return
}

/*
一个桶能装多少水，取决于左边木板最大高度和右边木板最大高度。
法 2：

	分别计算所有位置的左侧最大值和右侧最大值；
	前后缀分解：
		前缀数组表示当前位置最左边到当前位置的最大高度；
		后缀数组表示当前位置最右边到当前位置的最大高度；
*/
func trap_2(height []int) int {
	left_max, right_max := 0, 0
	pre_max := make([]int, len(height))
	for i := 0; i < len(height); i++ {
		pre_max[i] = left_max
		left_max = max(left_max, height[i])
	}

	suf_max := make([]int, len(height))
	for j := len(height) - 1; j >= 0; j-- {
		suf_max[j] = right_max
		right_max = max(right_max, height[j])
	}

	ans := 0
	for k := 0; k < len(height); k++ {
		capacity := min(pre_max[k], suf_max[k])
		if height[k] < capacity {
			ans += (capacity - height[k])
		}
	}

	return ans
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

