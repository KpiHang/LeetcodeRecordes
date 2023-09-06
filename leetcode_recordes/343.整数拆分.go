/*
 * @lc app=leetcode.cn id=343 lang=golang
 *
 * [343] 整数拆分
 */

// @lc code=start
// 代录；DP；
func integerBreak(n int) int {
	/**
		动态五部曲
		1.确定dp下标及其含义
		2.确定递推公式
		3.确定dp初始化
		4.确定遍历顺序
		5.打印dp
	**/
	dp := make([]int, n+1)
	dp[1], dp[2] = 0, 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ { // 成绩最大一定是分成相近的两个数，两个数的情况就是平均分；
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

