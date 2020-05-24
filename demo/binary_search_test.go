package main

import (
	"testing"
)

// go test -v -test.run Test_BinarySearch
func Test_BinarySearch(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(BinarySearch(s, 5))
}
