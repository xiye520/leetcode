package demo

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

func TestBubbleSort(t *testing.T) {
	array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	BubbleSort(array)
	fmt.Println("after sort array:", array)
}

func TestSelectSort(t *testing.T) {
	array := []int{12, 84, 5, 63, 37, 9, 56, 1399}
	fmt.Println("befor sort array:", array)
	SelectSort(array)
	fmt.Println("after sort array:", array)
}

func QuickSort(arr []int) []int {
	innerQuickSort(arr, 0, len(arr)-1)
	return arr
}

func innerQuickSort(arr []int, lo, hi int) {
	fmt.Println("lo, hi", lo, hi)
	if lo >= hi {
		return
	}
	mid := innerParrtition(arr, lo, hi)
	innerQuickSort(arr, lo, mid)
	innerQuickSort(arr, mid+1, hi)
}

func innerParrtition(arr []int, lo, hi int) int {
	i, j := lo+1, hi
	k := arr[lo]
	for {
		for i < hi && k >= arr[i] {
			i++
		}
		for j > lo && k <= arr[j] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i] // 交换
	}
	arr[lo], arr[j] = arr[j], arr[lo]
	return j
}
