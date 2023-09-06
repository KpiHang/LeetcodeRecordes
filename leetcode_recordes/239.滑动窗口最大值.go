/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
package main

type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Front() int {
	return m.queue[0]
}

func (m *MyQueue) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}
func (m *MyQueue) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}

// 下面是自己尝试的失败的方法
func maxSlidingWindow_Self_Failed(nums []int, k int) []int {
	if len(nums) == 1 && k == 1 {
		return []int{nums[0]}
	}
	// 初始化第一个窗口的最大值；假定为第一个位置；
	ans := []int{}
	curr_max, curr_index := nums[0], 0
	left, right := 0, 0
	for right < len(nums) {
		// curr_max = max(curr_max, nums[right])
		if curr_max < nums[right] {
			curr_max = nums[right]
			curr_index = right
		}
		if right-left+1 == k {
			ans = append(ans, curr_max)
			if curr_index == left {
				curr_max = nums[left+1]
			}
			left++
		}
		right++
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
