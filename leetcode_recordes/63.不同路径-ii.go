/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
// 代录;DP;先62再这个；Self的优化；
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 便于直接使用62题的代码；

	// 定义dp数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 { // 初始化行时遇到obstacle；后面就不用循环了，且默认值为0；
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 { // 初始化列时遇到obstacle；
			break
		}
		dp[0][j] = 1
	}

	// 遍历；
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果此点是故障点，默认为零值，所以不需要额外赋值0；
			// 不是故障点，则正常计算；
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

