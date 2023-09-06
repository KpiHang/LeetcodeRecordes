/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
package main

// 广度优先搜索，借用队列
type position struct {
	rowIndex int
	colIndex int
}

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			bfs_one_island(grid, position{i, j})
			count++
		}
	}
	return count
}

func bfs_one_island(grid [][]byte, p position) {
	Queue := make([]position, 0)
	Queue = append(Queue, p)
	for len(Queue) > 0 {
		// 出队列
		curr := Queue[0]
		Queue = Queue[1:] // 出栈，先进后出；
		// 位置合理且值为1时；
		// - 位置合理，当curr在边界时上下左右，可能出界；
		if curr.rowIndex < len(grid) && curr.rowIndex >= 0 && curr.colIndex < len(grid[0]) && curr.colIndex >= 0 && grid[curr.rowIndex][curr.colIndex] == '1' {
			grid[curr.rowIndex][curr.colIndex] = '0'
			p_up := position{curr.rowIndex - 1, curr.colIndex}
			p_down := position{curr.rowIndex + 1, curr.colIndex}
			p_left := position{curr.rowIndex, curr.colIndex - 1}
			p_right := position{curr.rowIndex, curr.colIndex + 1}
			Queue = append(Queue, p_up, p_down, p_left, p_right)
		}

	}
}

// @lc code=end
