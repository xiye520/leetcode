package medium

import "math"

/*
695. 岛屿的最大面积
给定一个包含了一些 0 和 1 的非空二维数组 grid 。

一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
*/
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var ret int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				ret = int(math.Max(float64(ret), float64(dfs2(grid, r, c))))
			}

		}
	}

	return ret
}

func dfs2(grid [][]int, r, c int) int {
	if !(r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])) {
		return 0
	}

	if grid[r][c] != 1 {
		return 0
	}

	grid[r][c] = 0
	return dfs2(grid, r+1, c) + dfs2(grid, r-1, c) + dfs2(grid, r, c+1) + dfs2(grid, r, c-1) + 1
}
