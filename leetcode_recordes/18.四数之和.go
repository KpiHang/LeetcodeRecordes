/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

/*
相向双指针1 作业2
此法，总结：
- 作用在有序列表，一左一右，相向而行；
- 因为有序，大小确定信息。

拓展，三数之和：
- fixed + min1 + min2 > target; break;
- fixed + max1 + max2 < target; 下一个fixed congtinue;
*/
// @lc code=start
func fourSum(nums []int, target int) (ans [][]int) {
	var (
		n = len(nums)
	)
	sort.Ints(nums)

	// for i, x := range nums[:n-3] { // 这种写法 下标是从零开始的，具体看下面go的循环；
	for i := 0; i < n-3; i++ {
		x := nums[i]
		// 跳过相同数字
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2]+nums[i+3] > target { // 优化一；因为数组有序
			break
		}
		if x+nums[n-1]+nums[n-2]+nums[n-3] < target { // 优化二；因为数组有序
			continue
		}
		// for j, y := range nums[i+1 : n-2] {  这种写法 j 是从0开始； 此法错误；
		for j := i + 1; j < n-2; j++ { // 注意这两种方法区别；
			y := nums[j]
			// 跳过相同数字
			if j > i+1 && y == nums[j-1] {
				continue
			}
			if x+y+nums[j+1]+nums[j+2] > target { // 优化一；因为数组有序
				break
			}
			if x+y+nums[n-1]+nums[n-2] < target { // 优化二；因为数组有序
				continue
			}
			left, right := j+1, n-1
			for left < right {
				sum := x + y + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else { // !!! 只有和target相等的时候，才需要左右跳过
					ans = append(ans, []int{x, y, nums[left], nums[right]})
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					} // 跳过右侧相同的
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					} // 跳过左侧相同的
				}
			}
		}
	}
	return
}

// @lc code=end

