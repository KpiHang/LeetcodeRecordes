/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
// 代录；0-1bag;
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if abs(target) > sum {
		return 0
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	// 计算目标背包大小；
	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, num := range nums {
		for i := bag; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[bag]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

