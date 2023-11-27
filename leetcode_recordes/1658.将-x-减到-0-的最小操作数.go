/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start

// 方法二：滑动窗口
// https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/solutions/2047253/jiang-x-jian-dao-0-de-zui-xiao-cao-zuo-s-hl7u/
func minOperations(nums []int, x int) int {
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x { // 所有原始之和<x, 无法减到0
		return -1
	}

	shortest := n + 1
	right := 0
	lsum, rsum := 0, sum
	for left := -1; left < n; left++ {
		if left != -1 {
			lsum += nums[left]
		}
		for right < n && lsum+rsum > x {
			rsum -= nums[right]
			right++
		}
		if lsum+rsum == x {
			shortest = min(shortest, n-right+left+1)
		}
	}
	if shortest == n+1 {
		return -1
	}
	return shortest
}

// @lc code=end

