/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start

// 贪心算法；
// https://www.programmercarl.com/0455.%E5%88%86%E5%8F%91%E9%A5%BC%E5%B9%B2.html
func findContentChildren(g []int, s []int) int {
	// g kids require; s cookie ;
	var result int
	sort.Ints(g)
	sort.Ints(s)
	index := len(s) - 1                // control s
	for i := len(g) - 1; i >= 0; i-- { // control g
		if index >= 0 && s[index] >= g[i] { // index>0 饼干数量多于人数的情况；
			result++
			index--
		}
	}
	return result
}

// @lc code=end

