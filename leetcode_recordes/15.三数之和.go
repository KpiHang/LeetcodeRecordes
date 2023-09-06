/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start

/*
灵神：相向双指针1；(已分类)
Background: 167
无序数组nums，找三元组 [nums[i], nums[j], nums[k]]
满足：i != j、i != k 且 j != k，nums[i] + nums[j] + nums[k] == 0

nums排序后，不妨设 i<j<k，
nums[j] + nums[k] = -nums[i]，枚举i（每个i固定一个值，就和167题相似；）
*/
func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)

	for i, x := range nums[:n-2] { // n-2，n-1 留给j，k， i < j < k
		if i > 0 && x == nums[i-1] { // 跳过重复数字，eg [-1,0,1,2,-1,-4] -1 重复；
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 { // 优化1 最小的三个数相加都大于0，直接结束了；
			break
		}
		if x+nums[n-1]+nums[n-2] < 0 { // 优化2 x和最大两个数向加，小于0，x直接下一个了。
			continue
		}
		j, k := i+1, n-1 // i < j < k
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				} // 跳过左侧重复数字
				for k--; k > j && nums[k] == nums[k+1]; k-- {
				} // 跳过右侧重复数字

			}
		}
	}
	return
}

// @lc code=end

