/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
// SELF: 依旧是双队列层序遍历二叉树+数值操作的处理
func findBottomLeftValue(root *TreeNode) int {

	Q := []*TreeNode{root}
	var ans int
	for len(Q) > 0 {
		P := []*TreeNode{}
		ans = Q[0].Val
		for i := 0; i < len(Q); i++ {
			node := Q[i]
			if node.Left != nil {
				P = append(P, node.Left)
			}
			if node.Right != nil {
				P = append(P, node.Right)
			}
		}
		Q = P
	}
	return ans
}

// @lc code=end

