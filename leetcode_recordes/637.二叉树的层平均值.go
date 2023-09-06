/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
// Self: 双队列层序遍历模板+Val处理;-------秒杀------
func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		P := []*TreeNode{}
		sum := 0.0
		for i := 0; i < len(Q); i++ {
			node := Q[i]
			sum += float64(node.Val)
			if node.Left != nil {
				P = append(P, node.Left)
			}
			if node.Right != nil {
				P = append(P, node.Right)
			}
		}
		ans = append(ans, sum/float64(len(Q)))
		Q = P
	}
	return ans
}

// @lc code=end

