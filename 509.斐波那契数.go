/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// 0 1 1 2

// 代录；动态规划DP基础题；
func fib(n int) int {
	if n < 2 {
		return n
	}
	a, b, c := 0, 1, 0 // dp初始化
	for i := 1; i < n; i++ {
		c = a + b // 递归公式
		a, b = b, c
	}
	return c
}

func fib_self2(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib_self(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dp := []int{}
	dp = append(dp, 0)
	dp = append(dp, 1)
	for i := 2; i <= n; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}

// @lc code=end

