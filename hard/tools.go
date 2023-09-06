package hard

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

type ListNode struct {
	val  int
	Next *ListNode
}
