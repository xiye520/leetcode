package main

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := maxArea(height)
	t.Log("max area", area)
}

func TestHeapSort(t *testing.T) {
	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	// var inputList = []int{}
	// inputList = CreateList(inputList, 10)

	HeapSort(inputList)
	t.Log(inputList)
}

func TestMergeSort(t *testing.T) {
	// var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	var inputList = []int{}
	inputList = CreateList(inputList, 10)

	MergeSort(inputList, 0, len(inputList)-1)
	t.Log(inputList)
}

func Test_maxArea(t *testing.T) {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Fatalf("maxArea not match 49,it is: %d", result)
	}
	t.Logf("maxArea: %d", result)
}
