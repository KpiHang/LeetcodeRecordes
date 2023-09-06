/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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

// 参考：二叉树后序递归遍历
func isBalanced(root *TreeNode) bool {
	// 定义一个匿名函数；
	var get_height func(node *TreeNode) int
	get_height = func(node *TreeNode) int {
		if node == nil {
			return 0 // 叶子结点高为0
		}
		// 左
		left_height := get_height(node.Left)
		if left_height == -1 {
			return -1 // 说明左子树上有不平衡
		}

		// 右
		right_height := get_height(node.Right)
		if right_height == -1 {
			return -1 // 说明左子树上有不平衡
		}

		// 中
		if abs(left_height-right_height) > 1 {
			return -1
		} else {
			return 1 + max(left_height, right_height)
		}
	}

	if get_height(root) == -1 {
		return false
	} else {
		return true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

