/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
// 深度优先遍历，回溯法本质就是dfs；
var (
	path []int
	res  [][]int
)

func subsets(nums []int) [][]int {
	path = []int{}
	res = [][]int{}
	dfs(nums, 0)
	return res
}

// 回溯算法其实本质就是深度优先搜索dfs；
func dfs(nums []int, start int) {
	tmp := make([]int, len(path)) // 切片本质是对底层数组的反映；
	copy(tmp, path)
	res = append(res, tmp)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		dfs(nums, i+1)
		path = path[:len(path)-1]
	}
}

// @lc code=end

