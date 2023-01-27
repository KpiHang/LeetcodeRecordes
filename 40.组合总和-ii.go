/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
// Self&&code with backtrack templete; 排序避免重复;
var res [][]int
var used []bool

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	used = make([]bool, len(candidates))
	sort.Ints(candidates)
	backtrack(0, 0, target, candidates, []int{})
	return res
}

func backtrack(sum, stratIndex, target int, candidates, trace []int) {
	if sum == target {
		temp := make([]int, len(trace))
		copy(temp, trace)
		res = append(res, temp)
		return
	}
	if sum > target {
		return
	}
	for i := stratIndex; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == false {
			// candidates[i] == candidates[i-1] used[i - 1] == true，说明同一树*枝*candidates[i - 1]使用过。
			// candidates[i] == candidates[i-1] used[i - 1] == false，说明同一树*层*candidates[i - 1]使用过，
			// 要对同一树层使用过的元素进行跳过。
			continue
		}
		trace = append(trace, candidates[i])
		sum += candidates[i]
		used[i] = true
		backtrack(sum, i+1, target, candidates, trace)
		trace = trace[:len(trace)-1]
		sum -= candidates[i]
		used[i] = false
	}
}

// @lc code=end

