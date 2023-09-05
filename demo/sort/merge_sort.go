// 1.归并排序
package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// 产生n个随机数
func CreateList(list []int, n int) []int {
	s1 := rand.NewSource(time.Now().Unix())
	r1 := rand.New(s1)
	for i := 0; i < n; i++ {
		list = append(list, r1.Intn(100))
	}
	return list
}

func Merge(list []int, a, b int) {
	var (
		i, j int
		mid  int
		pos  int
	)

	temp := make([]int, len(list))
	copy(temp, list)
	i = a
	mid = (a + b) / 2
	j = mid + 1
	pos = a
	for i <= mid && j <= b {
		if temp[i] < temp[j] {
			list[pos] = temp[i]
			pos++
			i++
		} else {
			list[pos] = temp[j]
			pos++
			j++
		}
	}
	for i <= mid {
		list[pos] = temp[i]
		pos++
		i++
	}
	for j <= b {
		list[pos] = temp[j]
		pos++
		j++
	}
}

// 左闭右闭[a,b]
func MergeSort(list []int, a, b int) {
	if a+1 >= b {
		if list[a] > list[b] {
			list[a], list[b] = list[b], list[a]
		}
		return
	}
	// 排序包括下标(a+b)/2的值
	MergeSort(list, a, (a+b)/2)

	MergeSort(list, (a+b)/2+1, b)

	Merge(list, a, b)
}

func testMergeSort() {
	// var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	var inputList = []int{}
	inputList = CreateList(inputList, 10)

	MergeSort(inputList, 0, len(inputList)-1)
	fmt.Println(inputList)
}
