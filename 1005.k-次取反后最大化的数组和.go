/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
// 暴力遍历, self
// 此题的暴力法思路其实就是贪心算法，
// 局部最优就是把绝对值最大的复数变为整数，具体，排序然后最小的复数变为正数；
import "sort"

func largestSumAfterKNegations_1(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if k == 0 {
			return nums_sum(nums)
		}
		if nums[i] <= 0 && k != 0 {
			nums[i] = 0 - nums[i]
			k--
		}
		if nums[i] > 0 && k == 0 {
			break
		}
	}
	sort.Ints(nums)
	for i := 0; i < k; i++ {
		nums[0] = 0 - nums[0]
	}
	return nums_sum(nums)
}

func nums_sum(nums []int) (sum int) {
	for _, val := range nums {
		sum += val
	}
	return
}

// 贪心算法 + 按照绝对值大小排序；
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool { // 把绝对值大的排前面；
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 { // 两个k 负负得正；
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// @lc code=end

