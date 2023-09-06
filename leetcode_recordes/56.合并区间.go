/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
// self; 贪心；
func merge(intervals [][]int) (res [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][1] < intervals[i][0] {
			temp := make([]int, 2)
			copy(temp, intervals[i-1])
			res = append(res, temp)
		} else {
			intervals[i][0] = intervals[i-1][0]
			intervals[i][1] = max(intervals[i-1][1], intervals[i][1])
		}
	}
	return append(res, intervals[len(intervals)-1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end

