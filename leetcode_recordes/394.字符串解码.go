/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
package main

import (
	"math"
	"strings"
)

func decodeString(s string) string {
	num_stack, alpha_stack := []int{}, []string{}
	repeat_alphas := false // 左括号之后就开始吧入字母栈
	ans := ""              // 保存返回结果
	for _, char_uint8 := range s {
		if char_uint8 <= '9' && char_uint8 >= '0' {
			num_stack = append(num_stack, int(char_uint8-'0'))
		} else if char_uint8 == '[' {
			repeat_alphas = true
			continue // [ 不入栈
		} else if char_uint8 == ']' {
			repeat_alphas = false
			k := 0 // 重复次数的int
			for i := 0; i < len(num_stack); i++ {
				k = k + num_stack[i]*int(math.Pow10(len(num_stack)-1-i))
			}
			num_stack = []int{} // 出栈
			str_tmp := ""
			for len(alpha_stack) > 0 {
				str_tmp = alpha_stack[len(alpha_stack)-1] + str_tmp
				alpha_stack = alpha_stack[:len(alpha_stack)-1]
			}
			ans = ans + strings.Repeat(str_tmp, k)
			alpha_stack = []string{}
		}

		if repeat_alphas { // true 表示重复的字母入栈处理；
			alpha_stack = append(alpha_stack, string(char_uint8))
		}
	}
	return ans
}

// @lc code=end
