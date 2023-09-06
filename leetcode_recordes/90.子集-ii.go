/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
// dfs
var (
	res  [][]int
	path []int
)

func subsetsWithDup(nums []int) [][]int {
	res = [][]int{}
	path = []int{}
	sort.Ints(nums)
	dfs(nums, 0)
	return res
}

func dfs(nums []int, start int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		if i != start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

