/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
// Self: 二叉搜索树的中序遍历顺序是递增的；
func findMode(root *TreeNode) []int {
	var traversal func(node *TreeNode) []int
	traversal = func(node *TreeNode) []int {
		res := []int{}
		if node == nil {
			return res
		}

		left := traversal(node.Left)
		right := traversal(node.Right)

		res = append(res, left...)
		res = append(res, node.Val)
		res = append(res, right...)
		return res
	}
	res := traversal(root)
	ans := []int{}
	max := 0
	left, right := 0, 0
	for left <= right {
		for res[right] == res[left] {
			right++
		}
		if (right - left) < max {
			left = right
			continue
		}
		if (right - left) > max {
			ans = []int{}
			max = right - left
		}
		ans = append(ans, res[left])
		left = right
	}
	return ans
}

// @lc code=end

