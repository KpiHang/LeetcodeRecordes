/*
 * @lc app=leetcode.cn id=452 lang=golang
 *
 * [452] 用最少数量的箭引爆气球
 */

// @lc code=start
// self：贪心；一箭多射几个；
func findMinArrowShots(points [][]int) (ans int) {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	linked := points[0]
	for i := 1; i < len(points); i++ {
		if points[i][0] <= linked[1] {
			linked[0] = points[i][0]
			if points[i][1] < linked[1] {
				linked[1] = points[i][1]
			}
			continue
		}
		linked = points[i]
		ans++
	}
	return ans + 1
}

// @lc code=end

