/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// Self: 秒杀; 双队列层序遍历二叉树 + 节点Next域操作处理；
// 此题完全适配116题的代码！
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	Q := []*Node{root}
	for len(Q) > 0 {
		P := []*Node{}
		for i := 0; i < len(Q); i++ {
			curr := Q[i]
			if i < len(Q)-1 {
				curr.Next = Q[i+1]
			} else {
				curr.Next = nil
			}
			if curr.Left != nil {
				P = append(P, curr.Left)
			}
			if curr.Right != nil {
				P = append(P, curr.Right)
			}
		}
		Q = P
	}
	return root
}

// @lc code=end

