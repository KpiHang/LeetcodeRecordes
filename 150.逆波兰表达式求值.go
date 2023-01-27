/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
package main

import "strconv"

// Self：栈
func evalRPN(tokens []string) int {
	stack := []int{}
	opts := map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}
	for _, token := range tokens {
		if _, ok := opts[token]; ok {
			// ok就是此时是运算符，弹栈两个计算结果再入栈；
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			total := 0
			switch token {
			case "+":
				total = a + b
			case "-":
				total = a - b
			case "*":
				total = a * b
			case "/":
				total = a / b
			}
			stack = append(stack, total)
		} else {
			Int, _ := strconv.Atoi(token)
			stack = append(stack, Int)
		}
	}
	return stack[0]
}

// @lc code=end
