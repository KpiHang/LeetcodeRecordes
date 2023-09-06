/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	minux_flag := false
	if n == 0 || x == 1.0 {
		return 1.0
	}
	if n < 0 {
		minux_flag = true
		n *= -1
	}
	var mult float64 = 1.0
	for i := 0; i < n; i++ {
		mult *= x
	}
	if minux_flag {
		return 1 / mult
	}
	return mult
}

// @lc code=end

