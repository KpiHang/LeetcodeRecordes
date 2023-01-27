/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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
// self
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == low {
		root.Left = nil
	}
	if root.Val == high {
		root.Right = nil
	}

	if root.Val < low {
		root = trimBST(root.Right, low, high)
	}
	if root.Val > high {
		root = trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// @lc code=end

