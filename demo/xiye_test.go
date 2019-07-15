package demo

import "testing"

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	t.Log("max area", area)
}
