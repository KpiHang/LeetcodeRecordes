/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start

/*
灵神（灵茶山艾府）：相向双指针1
复杂度另外一种方式是考虑获取信息量的多少；
如果On的方法，获取的信息量和On方相等，那么复杂度不就降下来吗。
*/
func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for {
		s := numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1}
		} else if s > target { // 最小的 + 最大的 > target 说明 次小的+最大的 也>target
			right-- // 所以right--， 最小的 + 次小的 看看什么情况；
		} else {
			left++
		}
	}
}

/*
花费了O1 的是时间，知道了On的信息；（因为，有序数组）
*/

// @lc code=end

