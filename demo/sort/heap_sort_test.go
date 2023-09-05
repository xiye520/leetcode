package sort

import (
	"testing"
)

func Test_heap(t *testing.T) {
	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	// var inputList = []int{}
	// inputList = CreateList(inputList, 10)

	HeapSort(inputList)
	t.Log(inputList)
}
