/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Self: 中序递归遍历
func isValidBST(root *TreeNode) bool {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		res := []int{}
		if node == nil {
			return res
		}
		left := traversal(node.Left)
		right := traversal(node.Right)

		res = append(res, left...)  // L
		res = append(res, node.Val) // N
		res = append(res, right...) // R
		return res
	}

	ans := traversal(root)
	for i := 1; i < len(ans); i++ {
		if ans[i] <= ans[i-1] {
			return false
		}
	}
	return true
}

// @lc code=end

