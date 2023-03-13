/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
// 暴力法；但遍历一个整数的每一位数字需要背；
func monotoneIncreasingDigits(n int) int {
	for n > 0 {
		if increasingable(n) {
			return n
		}
		n--
	}
	return -1

}

func increasingable(n int) bool {
	digits := []int{}
	for n > 0 {
		digit := n % 10                // 取出最后一位数字
		digits = append(digits, digit) // 输出数字
		n /= 10                        // 去除最后一位数字
	}
	for i := 1; i < len(digits); i++ {
		if digits[i-1] < digits[i] {
			return false
		}
	}
	return true
}

// @lc code=end

