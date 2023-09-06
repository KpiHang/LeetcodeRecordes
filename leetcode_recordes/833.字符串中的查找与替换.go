/*
 * @lc app=leetcode.cn id=833 lang=golang
 *
 * [833] 字符串中的查找与替换
 */

// @lc code=start
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	// 拖太久了，直接抄；
	n, m := len(s), len(indices)

	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		ops[i] = []int{i, indices[i]}
	}
	sort.Slice(ops, func(i int, j int) bool {
		return ops[i][1] < ops[j][1]
	})

	ans := ""
	pt := 0
	for i := 0; i < n; {
		for pt < m && indices[ops[pt][0]] < i {
			pt++
		}
		succeed := false
		for pt < m && indices[ops[pt][0]] == i {
			if s[i:i+min(len(sources[ops[pt][0]]), n-i)] == sources[ops[pt][0]] {
				succeed = true
				break
			}
			pt++
		}
		if succeed {
			ans += targets[ops[pt][0]]
			i += len(sources[ops[pt][0]])
		} else {
			ans += string(s[i])
			i++
		}
	}
	return ans
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func findReplaceString_temp(s string, indices []int, sources []string, targets []string) string {
	// 没必要用list，直接用string即可，方法相同。

	// 使用第一个切片的排序变化，对其他切片进行相同的排序
	sort.Slice(sources, func(i, j int) bool {
		return indices[i] < indices[j]
	})
	sort.Slice(targets, func(i, j int) bool {
		return indices[i] < indices[j]
	})
	sort.Slice(indices, func(i, j int) bool {
		return indices[i] < indices[j]
	})

	s_list := make([]string, 0)
	cur := 0
	for i := 0; i < len(indices); i++ {
		s_list = append(s_list, s[cur:indices[i]])
		if s[indices[i]:indices[i]+len(sources[i])] != sources[i] {
			s_list = append(s_list, s[indices[i]:indices[i]+len(sources[i])])
		} else {
			s_list = append(s_list, targets[i])
		}
		cur = indices[i] + len(sources[i])
	}
	s_list = append(s_list, s[cur:])

	return strings.Join(s_list, "")
}

// @lc code=end

