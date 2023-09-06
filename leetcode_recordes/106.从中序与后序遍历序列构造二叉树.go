/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
// Self: 递归法（分治），Find的效率不高
func buildTree(inorder []int, postorder []int) *TreeNode {
	ans := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return ans
	}
	i, _ := Find(inorder, postorder[len(postorder)-1])
	left_inorder := inorder[:i]
	right_inorder := inorder[i+1:]
	left_postorder := postorder[:len(left_inorder)]
	right_postorder := postorder[len(left_inorder) : len(postorder)-1]

	if len(left_inorder) > 0 {
		ans.Left = buildTree(left_inorder, left_postorder)
	}
	if len(right_inorder) > 0 {
		ans.Right = buildTree(right_inorder, right_postorder)
	}
	return ans
}

// Find查找切片中的元素下标，时间效率低！
func Find(ls []int, val int) (int, bool) {
	for i, item := range ls {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// @lc code=end

