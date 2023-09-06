/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
// Self: 依旧是双队列层序遍历 + 数值处理操作模板；
func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	Q := []*TreeNode{root}
	for len(Q) > 0 {
		P := []*TreeNode{}
		max_level := ^(int(^uint32(0) >> 1)) // 此题限制的Val的最小值；
		for i := 0; i < len(Q); i++ {
			curr := Q[i]
			if curr.Val > max_level {
				max_level = curr.Val
			}
			if curr.Left != nil {
				P = append(P, curr.Left)
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
			}
		}
		res = append(res, max_level)
		Q = P
	}
	return
}

// @lc code=end

