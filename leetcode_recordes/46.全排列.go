/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
// dfs
var (
	path []int
	res  [][]int
	used map[int]bool
)

func permute(nums []int) [][]int {
	path = []int{}
	res = [][]int{}
	used = make(map[int]bool)
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
		if used[nums[i]] {
			continue
		}
		path = append(path, nums[i])
		used[nums[i]] = true
		dfs(nums)
		path = path[:len(path)-1]
		used[nums[i]] = false
	}
}

// @lc code=end

