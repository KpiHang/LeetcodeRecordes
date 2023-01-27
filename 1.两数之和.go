/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for index, value := range nums {
		if p, ok := hashmap[target-value]; ok {
			return []int{p, index}
		} else {
			hashmap[value] = index
		}
	}
	return nil
}

// @lc code=end

