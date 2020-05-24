package main

import (
	"testing"
)

func Test_MergeSort(t *testing.T) {
	//testMergeSort()
	// var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	var inputList = []int{}
	inputList = CreateList(inputList, 10)

	MergeSort(inputList, 0, len(inputList)-1)
	t.Log(inputList)
}
