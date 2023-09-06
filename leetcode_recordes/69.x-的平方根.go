/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
package main

// 二分查找：求x的平方根的整数部；
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left, right := 0, x/2
	for left <= right {
		mid := (left + right) / 2
		if mid*mid < x {
			left = mid + 1
		} else if mid*mid > x {
			right = mid - 1
		} else {
			return mid
		}
	}
	// 上面循环结束后，right在left的左侧；
	return right
}

// @lc code=end
