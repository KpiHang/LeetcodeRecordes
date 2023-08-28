/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	// 巧妙的逆向思维,空间复杂度 O(1) 时间复杂度 O(m+n)
	p1, p2 := m-1, n-1
	cur := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[cur] = nums1[p1]
			p1--
		} else {
			nums1[cur] = nums2[p2]
			p2--
		}
		cur--
	}
	for p2 >= 0 {
		nums1[cur] = nums2[p2]
		p2--
		cur--
	}
}

func merge_1(nums1 []int, m int, nums2 []int, n int) {
	// 本题题干并未要求空间复杂度
	// 空间复杂度 O(m+n) 时间复杂度 O(m+n)
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if nums1[p1] <= nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}

	if p1 < m {
		sorted = append(sorted, nums1[p1:]...)
	}
	if p2 < n {
		sorted = append(sorted, nums2[p2:]...)
	}

	copy(nums1, sorted)
}

// @lc code=end

