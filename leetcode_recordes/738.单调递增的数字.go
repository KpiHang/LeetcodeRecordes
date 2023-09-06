/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
// 代录；贪心：高位-1，所有低位都变为99999....;
// Go语言中字符串与Go的相互转换；
func monotoneIncreasingDigits(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}

	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { // 前一位大于后一位，前一位减1，后面的全部拉满9999....；
			ss[i-1] -= 1
			for j := i; j < n; j++ {
				ss[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(ss))
	return res
}

// @lc code=end

