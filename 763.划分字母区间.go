/*
 * @lc app=leetcode.cn id=763 lang=golang
 *
 * [763] 划分字母区间
 */

// @lc code=start
// 代录；not 严格的贪心，但思路巧妙；
func partitionLabels(s string) (res []int) {
	var marks [26]int // 用数组的形式表示26个英文字母，及其对应下标；
	size, left, right := len(s), 0, 0
	for i := 0; i < size; i++ {
		marks[s[i]-'a'] = i
	}

	for i := 0; i < size; i++ {
		right = max(right, marks[s[i]-'a'])
		if right == i {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end

