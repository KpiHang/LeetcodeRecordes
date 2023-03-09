/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
// 代录，贪心；
func canCompleteCircuit(gas []int, cost []int) int {
	totalSum, currSum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		totalSum += rest
		currSum += rest
		if currSum < 0 {
			start, currSum = i+1, 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}

// @lc code=end

