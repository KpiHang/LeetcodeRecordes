/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
// Self: 双队列层序遍历二叉树 + 数值处理操作；
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	Q := []*TreeNode{root}
	length := len(Q)
	for length > 0 {
		P := []*TreeNode{}

		// 判断当前层是否对称；
		if !CurrLevelisSymmetric(Q) {
			return false
		}
		length = 0 // length 复位，length只是每一层有效元素的个数；
		for i := 0; i < len(Q); i++ {
			curr := Q[i]
			if curr.Left != nil {
				P = append(P, curr.Left)
				length++
			} else { // 添加的辅助节点，因此队列长度不更新；
				P = append(P, &TreeNode{Val: -101})
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
				length++
			} else { // 添加的辅助节点，因此队列长度不更新；
				P = append(P, &TreeNode{Val: -101})
			}
		}
		Q = P
	}
	return true
}

func CurrLevelisSymmetric(level []*TreeNode) bool {
	left, right := 0, len(level)-1
	for left < right {
		if level[left].Val != level[right].Val {
			return false
		}
		left++
		right--
	}
	return true
}

// @lc code=end

