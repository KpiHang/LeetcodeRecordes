/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start

/*
相向指针1：作业1
从nums中选3个数，使得这3个数的和最接近target
和 15 题十分相近；先15

"假定每组输入只存在恰好一个解。"
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32 // 最好的结果，三数之和；
	)

	// 根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}
	for i, x := range nums[:n-2] {
		// 保证和上一次枚举的元素不相等
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		// 使用双指针枚举 b 和 c
		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			// 如果和为 target 直接返回答案
			if s == target {
				return target
			}
			update(s)
			if s > target {
				// 如果和大于 target，移动到下一个不相等的元素
				// 大了，right右边k 要移动
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}

			} else {
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

