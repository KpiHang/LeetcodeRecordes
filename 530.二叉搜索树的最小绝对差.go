/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
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
// Self: 二叉树中序递归遍历
func getMinimumDifference(root *TreeNode) int {

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

	ans := traversal(root)
	var MinDiff int = getAbsMinus(ans[0], ans[1])
	for i := 2; i < len(ans); i++ {
		temp := getAbsMinus(ans[i], ans[i-1])
		if temp < MinDiff {
			MinDiff = temp
		}
	}
	return MinDiff
}
func getAbsMinus(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// @lc code=end

