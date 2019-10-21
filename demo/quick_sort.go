package main

import "fmt"

// 2.快排

// 左闭右闭[a,b]
func QuickSort(list []int) {
	var (
		midValue   int
		head, tail int
	)
	if len(list) <= 1 {
		return
	}
	head = 1
	midValue = list[0]
	tail = len(list) - 1
	for head < tail {
		for list[head] < midValue {
			head++
			if head >= tail {
				goto END
			}
		}
		for list[tail] > midValue {
			tail--
			if head >= tail {
				goto END
			}
		}
		list[head], list[tail] = list[tail], list[head]
		head++
		tail--
	}
END:
	if list[head] > midValue {
		list[0], list[head-1] = list[head-1], list[0]

		QuickSort(list[0 : head-1])
		QuickSort(list[head:])
	} else {
		list[0], list[head] = list[head], list[0]

		QuickSort(list[0:head])
		QuickSort(list[head+1:])
	}
}

func test_QuickSort(inputList []int) {

	// var inputList = []int{}
	// inputList = CreateList(inputList, 10)

	QuickSort(inputList)
	fmt.Println(inputList)
}

func main() {
	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	test_QuickSort(inputList)
	//test_quickSort(inputList)
}

/*
快速排序的递归描述如下：
对序列L排序时

如果L为空，则排序结果为空
否则，在L中任选一个元素作为pivot，然后递归的将L中不大于pivot的元素排序，将结果置于pivot左侧，同时递归地将所有大于pivot的元素排序，将结果置于pivot的右侧。
递归的基本快速排序算法实现如下:
*/

func quickSort(source []int, l, u int) {
	if l < u {
		m := partition(source, l, u)
		quickSort(source, l, m-1)
		quickSort(source, m, u)
	}
}

func partition(source []int, l, u int) int { //划分
	var (
		pivot = source[l]
		left  = l
		right = l + 1
	)
	for ; right < u; right++ {
		if source[right] <= pivot {
			left++
			source[left], source[right] = source[right], source[left]
		}
	}
	source[l], source[left] = source[left], source[l]
	return left + 1
}

func test_quickSort(s []int) {
	//s := []int{10, 6, 7, 4, 2, 5}
	quickSort(s, 0, len(s))
	fmt.Println(s)
}
