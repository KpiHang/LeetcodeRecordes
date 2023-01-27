/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
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
// Self：前序递归遍历二叉树；
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// N
	root1.Val += root2.Val
	// L
	root1.Left = mergeTrees(root1.Left, root2.Left)
	// R
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// @lc code=end

