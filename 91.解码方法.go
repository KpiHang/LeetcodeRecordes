/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	// dp 数组
	dp := make([]int, len(s))

	if s[0] == '0' {
		return 0
	}

	// 初始化
	dp[0] = 1

	// 确定遍历顺序和范围
	for i := 1; i < len(s); i++ {
		// step 1 考虑一位，也就是当前位i;
		if s[i] != '0' {
			dp[i] = dp[i-1]
		}
		// step 2 考虑两位;
		num := 10*(int(s[i-1]-'0')) + int(s[i]-'0')
		if 10 <= num && num <= 26 {
			if i == 1 {
				dp[i] += 1
			} else {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[len(s)-1]
}

// @lc code=end

