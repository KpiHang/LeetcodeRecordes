/*
 * @lc app=leetcode.cn id=332 lang=golang
 *
 * [332] 重新安排行程
 */

// @lc code=start

// https://www.programmercarl.com/0332.%E9%87%8D%E6%96%B0%E5%AE%89%E6%8E%92%E8%A1%8C%E7%A8%8B.html

type pair struct {
	target  string // 机场名；
	visited bool   // 是否访问过；
}

//sort.Sort 它排序任何实现了 sort.Interface 接口的数组。
// sort.Interface{Len(), Less(), Swap()}
// 实现这三个方法，也就实现了这个接口，实现了这个接口就可以用sort.Sort进行排序；
type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].target < p[j].target // 直接飞机场字母比大小；
}

func findItinerary(tickets [][]string) []string {
	result := []string{}
	// map[出发机场] pairs[pair1{目的地,是否被访问过},  pair2{目的地,是否被访问过},] 可能有同一个起点飞往多个机场的票；
	targets := make(map[string]pairs)
	for _, ticket := range tickets { // ticket 每一张机票，包括起点，终点；
		if targets[ticket[0]] == nil { // 初始map中没有某出发机场的记录；
			targets[ticket[0]] = make(pairs, 0) // 创建记录；
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	} // 得到targets 知道每一个机场可以飞的其他机场；
	for k, _ := range targets {
		sort.Sort(targets[k])
	} // targets 每一个机场可以飞的其他机场 且 按从小到大排列；

	result = append(result, "JFK") // 起点只能从JFK开始；
	var backtracking func() bool
	backtracking = func() bool {
		if len(tickets)+1 == len(result) { // 票数+1 就是 所有的机场数；
			return true // 三趟航班，找到了4个机场说明找到了结果
		}

		// 取出起飞航班对应的目的地 result[len(result)-1]:是拿到起飞机场；
		for _, pair := range targets[result[len(result)-1]] {
			// pair.visited == true 接着循环下一个，没飞过的同源目的机场；
			if pair.visited == false { // 因为已经排过序了，这里只需要看是否到过就好；
				result = append(result, pair.target)
				pair.visited = true
				if backtracking() { // backtracking 为true 说明已经找到了，就不需要再找了；
					return true
				}
				result = result[:len(result)-1]
				pair.visited = false
			}
		}

		// 循环完，还没找到，说明没有符合条件的
		return false
	}

	backtracking()

	return result
}

// @lc code=end

