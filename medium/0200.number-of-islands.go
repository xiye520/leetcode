package medium

/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。
*/

/*
每次走到一个是 1 的格子，就搜索整个岛屿。
网格可以看成是一个无向图的结构，每个格子和它上下左右的四个格子相邻。如果四个相邻的格子坐标合法，且是陆地，就可以继续搜索。
在深度优先搜索的时候要注意避免重复遍历。我们可以把已经遍历过的陆地改成 2，这样遇到 2 我们就知道已经遍历过这个格子了，不进行重复遍历。
每遇到一个陆地格子就进行深度优先搜索，最终搜索了几次就知道有几个岛屿。
*/
// dfs 深度优先遍历算法
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var count int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == '1' {
				dfs(grid, r, c)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, r, c int) {
	if !(r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])) {
		return
	}

	if grid[r][c] != '1' {
		return
	}

	grid[r][c] = '0'
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
