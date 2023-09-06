/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start

// 回溯法模板应用，
// 此题应该注意分清层次，不同层的不能在同一层循环；
var letterMap = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) <= 0 {
		return res
	}
	backtrack(0, digits, "")
	return res
}

func backtrack(Index int, digits, tempString string) {
	if len(digits) == len(tempString) {
		res = append(res, tempString)
		return
	}
	letters := letterMap[digits[Index]-'0']
	for _, letter := range letters {
		tempString = tempString + string(letter)
		backtrack(Index+1, digits, tempString)
		tempString = tempString[:len(tempString)-1]
	}

}

// @lc code=end

