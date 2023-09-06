/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start

// backTracking template
var res []string

func restoreIpAddresses(s string) []string {
	res = []string{}
	backTracking([]string{}, 0, s)
	return res
}

func backTracking(trace []string, startIndex int, s string) {
	if len(trace) == 4 { // 够四段就不再递归了；
		if startIndex == len(s) { // startIndex到最后了，4段确认已经划分好了；
			str := strings.Join(trace, ".")
			res = append(res, str)
		}
		return
	}

	for i := startIndex; i < len(s); i++ {
		if i != startIndex && s[startIndex] == '0' { // 前导零的判断；
			break
		}
		str := s[startIndex : i+1]
		num, _ := strconv.Atoi(str)
		if num >= 0 && num <= 255 {
			trace = append(trace, str) // 符合条件的就进入下一层；
			backTracking(trace, i+1, s)
			trace = trace[:len(trace)-1]
		} else { // 不满足条件，后面也绝不可能，剪枝操作；
			break
		}
	}
}

// @lc code=end

