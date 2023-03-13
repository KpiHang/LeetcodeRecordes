/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
// Self - 代录；贪心：高位-1，所有低位全部拉满（99999...）
func monotoneIncreasingDigits(n int) int {
	digits := []int{}
	for n > 0 {
		digit := n % 10
		digits = append(digits, digit) // 倒序了，就要从前往后；
		n /= 10
	}
	for i := 1; i < len(digits); i++ {
		if digits[i-1] < digits[i] {
			digits[i]--
			for j := i - 1; j >= 0; j-- {
				digits[j] = 9
			}
		}
	}
	var res int
	for i, val := range digits {
		res += val * int(math.Pow(10.0, float64(i)))
	}

	return res
}

// @lc code=end

