package medium

import (
	"testing"
)

func Test_search(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(search(s, 6))
}
