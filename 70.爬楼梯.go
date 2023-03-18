/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// Self；动态规划DP；
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1) // dp数组及下标的定义: 到达第i阶阶梯的方法数；
	dp[1], dp[2] = 1, 2    // dp数组初始化；
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 递推公式；
	}
	// fmt.Println(dp[1:]) // 打印dp数组；
	return dp[n]
}

// @lc code=end

