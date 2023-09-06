/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
// Self：递归法构建二叉树；
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index, val := Find(nums)
	res := &TreeNode{Val: val,
		Left:  constructMaximumBinaryTree(nums[:index]),
		Right: constructMaximumBinaryTree(nums[index+1:]),
	}
	return res
}

func Find(nums []int) (int, int) {
	index, max := 0, nums[0]
	for i, value := range nums {
		if value > max {
			max = value
			index = i
		}
	}
	return index, max
}

// @lc code=end

