/*
 * @lc app=leetcode.cn id=435 lang=golang
 *
 * [435] 无重叠区间
 */

// @lc code=start
// self；贪心：发生关联时，取小区间（是之后尽量不发生关联；）
func eraseOverlapIntervals(intervals [][]int) int {
	ans := 0

	// 排序；
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] { // 发生两个区间相关联了；
			ans++
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1]) // 记录要删，但不是真的删，且贪心，取小区间，让后面发生关联的可能性更小点；
		}
	}
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end

