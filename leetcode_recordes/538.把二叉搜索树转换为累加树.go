/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
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
// 反中序遍历；
func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var traversal func(node *TreeNode) *TreeNode
	traversal = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		// 右
		node.Right = traversal(node.Right)
		// 中
		sum += node.Val
		node.Val = sum
		// 左
		node.Left = traversal(node.Left)
		return node
	}
	return traversal(root)
}

// @lc code=end

