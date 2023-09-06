/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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

// Self: 二叉树层序遍历 + 子不为空的深度记录
func minDepth(root *TreeNode) (depth int) {
	if root == nil {
		return 0
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		P := []*TreeNode{}
		depth++ // depth位置要注意
		for _, curr := range Q {
			if curr.Left != nil {
				P = append(P, curr.Left)
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
			}
			if curr.Left == nil && curr.Right == nil {
				return depth
			}
		}
		Q = P
	}
	return depth
}

// @lc code=end

