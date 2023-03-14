/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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

// 代录；贪心：监控要放在叶子节点的父节点上；
const inf = math.MaxInt64 / 2

func minCameraCover(root *TreeNode) int {
	var dfs func(*TreeNode) (a, b, c int)
	dfs = func(node *TreeNode) (a, b, c int) {
		if node == nil {
			return inf, 0, 0
		}
		lefta, leftb, leftc := dfs(node.Left)
		righta, rightb, rightc := dfs(node.Right)
		a = leftc + rightc + 1
		b = min(a, min(lefta+rightb, righta+leftb))
		c = min(a, leftb+rightb)
		return
	}
	_, ans, _ := dfs(root)
	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// @lc code=end

