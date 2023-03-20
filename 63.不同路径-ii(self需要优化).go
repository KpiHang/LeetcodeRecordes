/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
// Self;DP;先62再这个；
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0]) // 便于直接使用62题的代码；

	// 定义dp数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 { // 初始化行时遇到obstacle；
			for i_ := i; i_ < m; i_++ {
				dp[i][0] = 0
			}
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 { // 初始化列时遇到obstacle；
			for j_ := j; j_ < n; j_++ {
				dp[0][j] = 0
			}
			break
		}
		dp[0][j] = 1
	}

	// 遍历；
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// @lc code=end

