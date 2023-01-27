/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package main

// 参考；学习二分查找元素左边界，二分查找元素右边界；
func searchRange(nums []int, target int) []int {
	// 目标值开始位置：为 target 的左侧边界
	start := searchLeftBoard(nums, target)
	// 如果开始位置越界 或 目标值不存在，直接返回
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	// 目标值结束位置：为 target 的右侧边界
	// 查找元素+1的左边界-1，即为右边界；
	end := searchLeftBoard(nums, target+1) - 1
	return []int{start, end}
}

// 二分查找元素的左边界
func searchLeftBoard(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return left
}

// @lc code=end
