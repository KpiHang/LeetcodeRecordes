/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	curr := root
	for curr != nil {
		if curr.Val == key {
			if curr.Left != nil {
				curr.Val = curr.Left.Val
				curr.Left = nil
				break
			}
			if curr.Right != {
				\
			}
		}
	}
}

// @lc code=end

