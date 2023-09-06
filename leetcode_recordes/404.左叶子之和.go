/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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

// 参考代码随想录，后续递归遍历；
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil { // 判断root是否为空；
		return 0
	}
	if root.Left == nil && root.Right == nil { // 递归到了叶子节点，但叶子不一定是左叶子，所以在下面判断；
		return 0
	}

	// 左
	left_value := sumOfLeftLeaves(root.Left)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { // 左孩子存在且为左叶子结点；
		left_value = root.Left.Val
	}

	// 右
	right_value := sumOfLeftLeaves(root.Right)

	// 中
	sum := left_value + right_value

	return sum
}

// @lc code=end

