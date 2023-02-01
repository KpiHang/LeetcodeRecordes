/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start

// Way 1: [0,2] + [2:]
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	ans := 1
	prediff := nums[1] - nums[0]
	if prediff != 0 {
		ans = 2
	}
	for i := 2; i < n; i++ {
		currdiff := nums[i] - nums[i-1]
		if currdiff > 0 && prediff <= 0 || currdiff < 0 && prediff >= 0 {
			ans++
			prediff = currdiff
		}
	}
	return ans
}

// @lc code=end

