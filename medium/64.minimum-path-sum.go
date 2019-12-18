package medium

/*
64. 最小路径和
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1,6],
  [1,5,1,3],
  [4,2,1,2]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	cols := len(grid)
	rows := len(grid[0])

	dp := make([][]int, cols)
	for i := 0; i < cols; i++ {
		dp[i] = make([]int, rows)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < cols; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < rows; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < cols; i++ {
		for j := 1; j < rows; j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}

	return dp[cols-1][rows-1]
}

func minPathSum2(grid [][]int) int {
	// 已经默认 m 和 n 大于 0 了
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}
