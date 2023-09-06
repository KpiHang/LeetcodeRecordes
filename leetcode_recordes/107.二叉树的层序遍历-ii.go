/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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

// Self: 先从上到下层序（102题），再reverse结果列表；
func levelOrderBottom(root *TreeNode) [][]int {
	res := levelOrder(root)
	reverse(res)
	return res
}
func reverse(a [][]int) {
	left, right := 0, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]
		left, right = left+1, right-1
	}
}
func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

// @lc code=end

