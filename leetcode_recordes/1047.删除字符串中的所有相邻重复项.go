/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
// Official answer
func removeDuplicates(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

// SELF：用string实现栈完成；
func removeDuplicates_method2(s string) string {
	stack_str := ""
	for _, c := range s {
		if len(stack_str) > 0 && string(stack_str[len(stack_str)-1]) == string(c) {
			stack_str = stack_str[:len(stack_str)-1]
			continue
		} else {
			stack_str = stack_str + string(c)
		}
	}
	return stack_str
}

// SELF：用slice实现栈完成；
func removeDuplicates_method1(s string) string {
	stack := []string{}
	ans := ""
	for _, c := range s {
		if len(stack) > 0 && stack[len(stack)-1] == string(c) {
			stack = stack[:len(stack)-1]
			continue
		} else {
			stack = append(stack, string(c))
		}
	}
	for _, ss := range stack {
		ans = ans + ss
	}
	return ans
}

// @lc code=end

