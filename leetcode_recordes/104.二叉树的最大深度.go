/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
// Self: 二叉树层序遍历 + 层序记录
func maxDepth(root *TreeNode) (levels int) {
	if root == nil {
		return 0
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		levels++
		P := []*TreeNode{}
		for _, curr := range Q {
			if curr.Left != nil {
				P = append(P, curr.Left)
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
			}
		}
		Q = P
	}
	return
}

// @lc code=end

