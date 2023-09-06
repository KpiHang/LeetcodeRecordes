/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start

// Self: 贪心；
// 我认为的贪心是，先把大钱找了，零钱尽可能多的留手里；

func lemonadeChange(bills []int) bool {
	changes := make(map[int]int, 0)
	for i := range bills {
		change := bills[i] - 5
		if _, ok := changes[bills[i]]; !ok {
			changes[bills[i]] = 1
		} else {
			changes[bills[i]] += 1
		}
		// 先把10元，大额找出去；留尽可能多的小额；
		if change >= 10 {
			if _, ok := changes[10]; ok && changes[10] != 0 {
				changes[10] -= 1
				change -= 10
			}
		}
		// 小额可以应对多种情况；
		for change > 0 { // change：0 or 5 or 15
			if _, ok := changes[5]; ok && changes[5] != 0 {
				changes[5] -= 1
				change -= 5
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

