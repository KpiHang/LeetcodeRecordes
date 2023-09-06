/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
// 代码随想录参考
// https://www.programmercarl.com/0051.N%E7%9A%87%E5%90%8E.html

func solveNQueens(n int) [][]string {
	var res [][]string

	// 创建并初始化棋盘......；
	chessboard := make([][]string, n)
	for i := 0; i < n; i++ {
		chessboard[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			chessboard[i][j] = "."
		}
	}

	var backtrack func(int)
	backtrack = func(row int) { // 标记当前行号，从0开始；
		if row == n { // row 从0开始，row==n说明结束了；
			temp := make([]string, n)
			for i, rowStr := range chessboard {
				temp[i] = strings.Join(rowStr, "")
			}
			res = append(res, temp)
			return
		}
		for i := 0; i < n; i++ {
			if isValid(n, row, i, chessboard) { // n， 目标位置的行，列号，棋盘；
				chessboard[row][i] = "Q"
				backtrack(row + 1) // 下一行；
				chessboard[row][i] = "."
			}
		}
	}

	backtrack(0)
	return res
}

func isValid(n, row, col int, chessboard [][]string) bool {
	// 方向1 正上方
	for i := row - 1; i >= 0; i-- {
		if chessboard[i][col] == "Q" {
			return false
		}
	}
	// 方向2 左斜上方
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}
	// 方向3 右斜上方
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == "Q" {
			return false
		}
	}

	return true
}

// @lc code=end

