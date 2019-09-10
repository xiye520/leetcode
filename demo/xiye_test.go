package demo

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
