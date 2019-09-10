package main

import "fmt"

func main() {
	a := []int{38, 100, 4, 2, 8, 80, 10}
	qsort(a, 0, len(a)-1)
	fmt.Println(a)
}

//快速排序  找到一个数,左边都比他小,右边都比他大
//38, 1, 4, 5, 10
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

func qsort(a []int, left, right int) {
	//如果左边等于右边,就是只有一个元素,不需要排序,直接返回
	if left >= right {
		return
	}
	mid := partion(a, left, right)
	qsort(a, left, mid-1)
	qsort(a, mid+1, right)
}
