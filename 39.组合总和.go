/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start

// self with backtrack templete; include used elem
var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = [][]int{}
	backtrack(0, target, 0, []int{}, candidates)
	return res
}

func backtrack(sum, target, Index int, trace, candidates []int) {
	if sum == target {
		temp := make([]int, len(trace))
		copy(temp, trace)
		res = append(res, temp)
		return
	}
	if sum > target {
		return
	}
	for i := Index; i < len(candidates); i++ {
		sum += candidates[i]
		trace = append(trace, candidates[i])
		backtrack(sum, target, i, trace, candidates)
		sum -= candidates[i]
		trace = trace[:len(trace)-1]
	}
}

// @lc code=end

