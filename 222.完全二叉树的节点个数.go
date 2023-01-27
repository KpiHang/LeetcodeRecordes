/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
// Self：依旧可以用 双队列层序遍历模版 + 每层节点个数相加；
func countNodes(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		P := []*TreeNode{}
		ans += len(Q)
		for i := 0; i < len(Q); i++ {
			curr := Q[i]
			if curr.Left != nil {
				P = append(P, curr.Left)
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
			}
		}
		Q = P
	}
	return ans
}

// @lc code=end

