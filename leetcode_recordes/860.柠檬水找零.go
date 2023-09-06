/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
// 代录；贪心；
// 思路和self 贪心是一样的，不过数据表示上没有我写的那么麻烦；
// 一共不超过三个，没必要用字典；
func lemonadeChange(bills []int) bool {
	ten, five := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 { // 5元 加一张，且不需要找零；
			five++
		} else if bills[i] == 10 { // 10元 ，找一张5元即可；
			if five == 0 {
				return false
			}
			ten++
			five--
		} else {
			if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

