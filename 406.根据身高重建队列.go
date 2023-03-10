/*
 * @lc app=leetcode.cn id=406 lang=golang
 *
 * [406] 根据身高重建队列
 */

// @lc code=start

// 代录；贪心：让身高高的先出现，那么k就比较好解决了；
func reconstructQueue(people [][]int) [][]int {
	// 先将身高从大到小排序，确定最大个子的位置；
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] // 当身高相同时，将K按照从小到大；
		}
		return people[i][0] > people[j][0] // 身高按照由大到小的顺序来排；
	})

	// 再按照K进行插入排序，优先插入k小的；// 这里不是插入排序；
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1]+1:], result[info[1]:]) // 把元素插入到下标为k的位置；
		result[info[1]] = info
	}
	return result
}

// @lc code=end

