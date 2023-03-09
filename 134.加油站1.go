/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
// 代录，不是严格的贪心算法，但思路很好！
func canCompleteCircuit(gas []int, cost []int) int {
	min, curSum := int(^uint(0)>>1), 0
	for i := 0; i < len(gas); i++ {
		rest := gas[i] - cost[i]
		curSum += rest
		if curSum < min {
			min = curSum
		}
	}
	if curSum < 0 { // case 1；
		return -1
	}
	if min >= 0 { // case 2;
		return 0
	}
	for j := len(gas) - 1; j >= 0; j-- { // case 3;
		min += (gas[j] - cost[j])
		if min >= 0 {
			return j
		}
	}
	return -1
}

// @lc code=end

