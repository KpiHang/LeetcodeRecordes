/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
// Self with backtrack templete
var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = [][]int{}
	if n < 6 || k <= 0 || k > n {
		return res
	}
	backtrack(k, n, 1, []int{})
	return res
}

func backtrack(k, n, start int, track []int) {
	if len(track) == k {
		if sum(track) == n {
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
		}
		return
	}

	// 剪枝↓↓↓
	if len(track)+9-start+1 < k {
		return
	}
	// 剪枝↑↑↑

	for i := start; i <= 9; i++ {
		track = append(track, i)
		backtrack(k, n, i+1, track)
		track = track[:len(track)-1]
	}
}

func sum(nums []int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

// @lc code=end

