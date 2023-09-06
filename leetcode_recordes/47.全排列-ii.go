/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
// dfs
var (
	path []int
	res  [][]int
	used []bool
)

func permuteUnique(nums []int) [][]int {
	path = []int{}
	res = [][]int{}
	used = make([]bool, len(nums))
	sort.Ints(nums)
	dfs(nums)
	return res
}

func dfs(nums []int) {
	if len(path) == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs(nums)
			path = path[:len(path)-1]
			used[i] = false
		}

	}
}

// @lc code=end

