/*
 * @lc app=leetcode.cn id=2605 lang=golang
 *
 * [2605] 从两个数字数组里生成最小数字
 */

// @lc code=start

/*
简单题，分类讨论；
case 1：nums1、nums2 中有相同元素，ans即为这个元素；
case 2：nums1、nums2 中没有相同元素
*/
func minNumber(nums1 []int, nums2 []int) int {
	// min1, min2 分别记录两个数组中的最小值
	min1, min2 := 10, 10
	same_min := 10 // 两个数组中相同元素的最小值

	// hash表记录nums1，nums2中如果有命中的，就是有相同元素 case1
	// 构造hash表；
	hash := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]] = i
		min1 = min(min1, nums1[i])
	}
	// 是否命中；
	for i := 0; i < len(nums2); i++ {
		// 命中，ans is the same element
		if _, ok := hash[nums2[i]]; ok {
			same_min = min(same_min, nums2[i])
		}
		min2 = min(min2, nums2[i])
	}

	if same_min < 10 {
		return same_min
	}

	return min(10*min1+min2, min1+10*min2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

