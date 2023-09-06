/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
// 回溯法模板; 剪枝;
var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	if n <= 0 || k <= 0 || k > n {
		return res
	}
	backtrack(n, k, 1, []int{})
	return res
}

func backtrack(n, k, start int, track []int) {
	if len(track) == k { // 递归终止条件；
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
		return
	}
	// 剪枝；
	if len(track)+n-start+1 < k {
		return
	}

	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1] // 回溯；
	}
}

// @lc code=end

