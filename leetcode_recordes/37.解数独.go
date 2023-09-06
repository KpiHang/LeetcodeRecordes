/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start

// 代码随想录
// https://www.programmercarl.com/0037.%E8%A7%A3%E6%95%B0%E7%8B%AC.html
func solveSudoku(board [][]byte) {
	var backtracking func() bool
	backtracking = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}

				// try 1-9
				for k := '1'; k <= '9'; k++ {
					if isValid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backtracking() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}
	backtracking()
}

func isValid(row, col int, k byte, board [][]byte) bool {
	// col whether repeat
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}

	// row
	for i := 0; i < 9; i++ {
		if board[i][col] == k {
			return false
		}
	}

	// 3*3
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3

	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}

// @lc code=end

