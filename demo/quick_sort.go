package main

import (
	"fmt"
)

func quickSort111(values []int, left int, right int) {
	key := values[left] //取出第一项
	p := left
	i, j := left, right

	for i <= j {
		//由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}

		//第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}

		if values[i] <= key && i <= p {
			i++
		}

		if i < p {
			values[p] = values[i]
			p = i
		}

		values[p] = key
		if p-left > 1 {
			quickSort111(values, left, p-1)
		}
		if right-p > 1 {
			quickSort111(values, p+1, right)
		}
	}

}

func QuickSort(values []int) {
	quickSort111(values, 0, len(values)-1)
}

// 快速排序方式二
func quickSort6(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort6(sortArray, left, i-1)
		quickSort6(sortArray, j+1, right)
	}
}

func QuickSort6(values []int) []int {
	quickSort6(values, 0, len(values)-1)
	return values
}

// -----------------------------------

// 快排方案1
func quickSort(list []int) {
	// 左闭右闭[a,b]
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

		quickSort(list[0 : head-1])
		quickSort(list[head:])
	} else {
		list[0], list[head] = list[head], list[0]

		quickSort(list[0:head])
		quickSort(list[head+1:])
	}
}

func test_quickSort(inputList []int) {

	// var inputList = []int{}
	// inputList = CreateList(inputList, 10)

	quickSort(inputList)
	fmt.Println(inputList)
}

// 快排方案2
func quickSort2(a []int, left, right int) {
	//如果左边等于右边,就是只有一个元素,不需要排序,直接返回
	if left >= right {
		return
	}
	mid := partion(a, left, right)
	quickSort2(a, left, mid-1)
	quickSort2(a, mid+1, right)
}

// 快速排序  找到一个数,左边都比他小,右边都比他大
// 38, 1, 4, 5, 10
func partion(a []int, i, j int) int {
	//var i ,j = left,right
	//1.当左边<右边
	for i < j {
		//2.1当右边大于左边 什么也不做
		for j > i && a[j] > a[i] {
			j--
		}
		//2.2 否则,交换位置
		a[j], a[i] = a[i], a[j]
		//2.3 左边都小于自己 什么也不做
		for i < j && a[i] < a[j] {
			i++
		}
		//2.4 否则交换位置
		a[j], a[i] = a[i], a[j]
	}
	//返回的是,一个数字的正确角标,然后可以根据这个角标去划分两个区域
	return i
}

func test_quickSort2(a []int) {
	//a := []int{38, 100, 4, 2, 8, 80, 10}
	quickSort2(a, 0, len(a)-1)
	fmt.Println(a)
}

// 快排方案3
func quickSort3(arr []int, l, u int) {
	/*
		快速排序
		快速排序的递归描述如下：
		对序列L排序时

		如果L为空，则排序结果为空
		否则，在L中任选一个元素作为pivot，然后递归的将L中不大于pivot的元素排序，将结果置于pivot左侧，同时递归地将所有大于pivot的元素排序，将结果置于pivot的右侧。
	*/
	if l < u {
		m := partition3(arr, l, u)
		quickSort3(arr, l, m-1)
		quickSort3(arr, m, u)
	}

}

func partition3(source []int, l, u int) int { //划分
	pivot := source[l]
	left := l
	right := l + 1

	for ; right < u; right++ {
		if source[right] <= pivot {
			left++
			source[left], source[right] = source[right], source[left]
		}
	}

	source[l], source[left] = source[left], source[l]
	return left + 1
}

func test_quickSort3(s []int) {
	//s := []int{10, 6, 7, 4, 2, 5}
	quickSort3(s, 0, len(s))
	fmt.Println(s)
}

// 快排方案4--双向划分
func quickSort4(source []int, l, u int) {
	/*
		在所有元素都相等的极端情况下， 双向划分可以将数组划分为两段长度相等的子数组， 虽然发生了n/2次不必要的交换，
		但因为总体划分是平衡的， 所以总体的时间复杂度还为: O(nlgn)
	*/
	if l < u-1 {
		m := partition4(source, l, u)
		quickSort4(source, l, m-1)
		quickSort4(source, m, u)
	}
}

func partition4(source []int, l, u int) int { //双向划分
	var (
		pivot = source[l]
		left  = l
		right = u
	)
	for {
		left++
		for left < u && source[left] < pivot {
			left++
		}
		right--
		for right >= l && source[right] > pivot {
			right--
		}
		if left >= right {
			break
		}
		source[left], source[right] = source[right], source[left]
	}
	source[l], source[right] = source[right], source[l]
	return right + 1
}

func test_quickSort4(s []int) {
	//s := []int{10, 6, 7, 4, 2, 5}
	quickSort4(s, 0, len(s))
	fmt.Println(s)
}

// 快排三路划分--简版
func quickSort5(source []int, l, u int) {
	var (
		i, j, k, pivot int
	)
	if l < u-1 {
		i = l
		j = u
		pivot = source[l]

		for k = l + 1; k < j; k++ {
			for pivot < source[k] {
				j--
				source[j], source[k] = source[k], source[j]
			}
			if source[k] < pivot {
				source[k], source[i] = source[i], source[k]
				i++
			}
		}
		quickSort5(source, l, i)
		quickSort5(source, k, u)
	}
}

// 快排三路划分--繁版
func quickSort51(source []int, l, u int) {
	var (
		i, j, p, q, pivot int
	)
	if l < u-1 {
		i, p = l, l
		j, q = u, u
		pivot = source[l]

		for {
			i++
			for i < u && source[i] < pivot {
				i++
			}
			j--
			for j >= l && source[j] > pivot {
				j--
			}
			if i >= j {
				break
			}
			source[i], source[j] = source[j], source[i]
			if source[i] == pivot {
				p++
				source[p], source[i] = source[i], source[p]
			}
			if source[j] == pivot {
				q--
				source[q], source[j] = source[j], source[q]
			}
		}
		if i == j && source[i] == pivot { //特殊情况
			j--
			i++
		}
		for k := l; k <= p; k++ {
			source[k], source[j] = source[j], source[k]
			j--
		}
		for k := u - 1; k >= q; k-- {
			source[k], source[i] = source[i], source[k]
			i++
		}
		quickSort51(source, l, j+1)
		quickSort51(source, i, u)
	}
}

func test_quickSort5(s []int) {
	//s := []int{10, 6, 7, 4, 2, 5}
	//quickSort5(s, 0, len(s))
	quickSort51(s, 0, len(s))
	fmt.Println(s)
}

//func main() {
//	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
//	//test_quickSort(inputList)
//	//test_quickSort2(inputList)
//	//test_quickSort3(inputList)
//	//test_quickSort4(inputList)
//	test_quickSort5(inputList)
//}
