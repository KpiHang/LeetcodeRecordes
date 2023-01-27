/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
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

// 参考：官方题解；
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Self：前序遍历
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	var traversal func(p *TreeNode, q *TreeNode) bool // 匿名函数 + 闭包完成遍历；
	traversal = func(p *TreeNode, q *TreeNode) bool {
		if p == nil && q == nil { // 同时为nil
			return true
		} else if p != nil && q != nil { // 同时不为nil
			if p.Val != q.Val {
				return false
			}
			if !traversal(p.Left, q.Left) || !traversal(p.Right, q.Right) { // 子树里已经有false了
				return false
			}
			return true
		} else { // 一个为nil，一个不为nil
			return false
		}
	}

	return traversal(p, q)
}

// @lc code=end

