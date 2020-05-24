package main

import (
	"testing"
)

func Test_maxArea(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Fatalf("maxArea not match 49,it is: %d", result)
	}
	t.Logf("maxArea: %d", result)

}
