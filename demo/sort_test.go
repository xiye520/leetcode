package main

import (
	"fmt"
	"testing"
)

//1.冒泡排序，向后移动，从小到大排序,其思想为相邻两个数进行比较，将较大的滞后，时间复杂度O(N^2
func BubbleSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

//2.选择排序，从后面找个最小的放在前面的位置，从小到大排序,时间复杂度O(N^2)
func SelectSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

//1.冒泡排序，向后移动，从小到大排序,其思想为相邻两个数进行比较，将较大的滞后，时间复杂度O(N^2
func BubbleSort1(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

//2.选择排序，从后面找个最小的放在前面的位置，从小到大排序,时间复杂度O(N^2)
func SelectSort1(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

func TestBubbleSort(t *testing.T) {
	array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	BubbleSort1(array)
	fmt.Println("after sort array:", array)
}

func TestSelectSort(t *testing.T) {
	array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	SelectSort1(array)
	fmt.Println("after sort array:", array)
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

func Test_quick_sort(t *testing.T) {
	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	//test_quickSort(inputList)
	//test_quickSort2(inputList)
	//test_quickSort3(inputList)
	//test_quickSort4(inputList)
	test_quickSort5(inputList)
}
