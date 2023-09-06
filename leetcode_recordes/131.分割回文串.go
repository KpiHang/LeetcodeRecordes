/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
// Backtrack templete
var res [][]string

func partition(s string) [][]string {
	res = [][]string{}
	BackTracking(0, []string{}, s)
	return res
}

func BackTracking(startIndex int, trace []string, s string) {
	if startIndex == len(s) {
		temp := make([]string, len(trace))
		copy(temp, trace)
		res = append(res, temp)
		return
	}
	for i := startIndex; i < len(s); i++ {
		if IsPalindromic(startIndex, i, s) {
			trace = append(trace, s[startIndex:i+1])

		} else {
			continue
		}
		BackTracking(i+1, trace, s)
		trace = trace[:len(trace)-1]
	}
}

func IsPalindromic(left, right int, s string) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

// @lc code=end

