/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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

// 树形DP；下标0：不偷；下标1：偷；
func rob(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	return max(robTree(root)[0], robTree(root)[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 下标0：不偷；下标1：偷；
func robTree(root *TreeNode) []int {
	if root == nil { // 叶子节点，偷和不偷都是0；
		return []int{0, 0}
	}
	// 后序遍历：左右后；
	left, right := robTree(root.Left), robTree(root.Right)
	robCur := left[0] + right[0] + root.Val
	notRobCur := max(left[1], left[0]) + max(right[1], right[0])
	return []int{notRobCur, robCur}
}

// @lc code=end

