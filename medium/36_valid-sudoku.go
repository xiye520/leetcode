package medium

/*
6. 有效的数独

判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的
*/

func isValidSudoku(board [][]byte) bool {
	var row [9][10]bool   // 行
	var cow [9][10]bool   // 列
	var block [9][10]bool // 块
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			var index = i/3*3 + j/3 //块坐标
			if board[i][j] != '.' {
				var tmp = board[i][j] - '0'
				if row[i][tmp] {
					return false
				} else {
					row[i][tmp] = true
				}
				if cow[j][tmp] {
					return false
				} else {
					cow[j][tmp] = true
				}
				if block[index][tmp] {
					return false
				} else {
					block[index][tmp] = true
				}

			}
		}

	}
	return true
}

func isValidSudoku2(board [][]byte) bool {
	n := len(board)
	rectSets := make([]int16, n*n/9) // sets for inner rectangles
	for i := 0; i < n; i++ {
		rowSet, colSet := int16(0), int16(0) // sets for a row and a column
		for j := 0; j < n; j++ {
			// checking a column and an inner rectangle
			if num := board[i][j]; num != '.' {
				numBit := int16(1 << (num - '0')) // calc a #num bit
				if rowSet&numBit > 0 {
					return false
				}
				rowSet |= numBit // set the bit

				// set the bit for a inner rectangle: ((n/3)*(i/3), j/3) - its coordinates
				if rectSets[(n/3)*(i/3)+j/3]&numBit > 0 {
					return false
				}
				rectSets[(n/3)*(i/3)+j/3] |= numBit
			}

			// checking column
			if num := board[j][i]; num != '.' {
				numBit := int16(1 << (num - '0'))
				if colSet&numBit > 0 {
					return false
				}
				colSet ^= numBit
			}
		}
	}
	return true
}
