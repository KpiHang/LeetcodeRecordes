/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
// 代录；0-1背包；
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 初始化dp[0][0]=0，默认值为0就不用初始化了；
	for i := 0; i < len(strs); i++ {
		zero, one := 0, 0
		for _, v := range strs[i] {
			if v == '0' {
				zero++
			}
		}
		one = len(strs[i]) - zero
		for j := m; j >= zero; j-- {
			for k := n; k >= one; k-- {
				dp[j][k] = max(dp[j][k], dp[j-zero][k-one]+1)
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

