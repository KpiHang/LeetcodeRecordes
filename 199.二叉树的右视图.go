/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
// Self: 双队列层序遍历模板+Val处理
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return ans
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		P := []*TreeNode{}
		for i := 0; i < len(Q); i++ {
			node := Q[i]
			if node.Left != nil {
				P = append(P, node.Left)
			}
			if node.Right != nil {
				P = append(P, node.Right)
			}
		}
		ans = append(ans, Q[len(Q)-1].Val)
		Q = P
	}
	return ans
}

// @lc code=end

