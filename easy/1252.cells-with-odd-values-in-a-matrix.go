package easy

func oddCells(n int, m int, indices [][]int) int {
	x, y := make([]int, n), make([]int, m)
	for _, in := range indices {
		x[in[0]]++ // 行
		y[in[1]]++ // 列
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans += (x[i] + y[j]) % 2
		}
	}

	return ans
}

func oddCells2(n int, m int, indices [][]int) int {
	var arr [][]int = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
	}

	for i := 0; i < len(indices); i++ {
		var row int = indices[i][0]
		var col int = indices[i][1]

		for j := 0; j < m; j++ {
			arr[row][j] += +1
		}
		for j := 0; j < n; j++ {
			arr[j][col] += 1
		}
	}

	var count int = 0
	for _, val := range arr {
		for _, val1 := range val {
			if val1%2 == 1 {
				count += 1
			}
		}
	}
	return count
}
