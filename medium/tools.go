package medium

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// abs 求绝对值，返回正整数
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
