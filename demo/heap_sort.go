/*
3.堆排序
推荐文章：
https://www.cnblogs.com/cheng...
https://blog.csdn.net/MoreWin...
*/
// i为处理节点,其是从孩子节点向下影响堆。
// n为堆的节点数
// 其存在一个前提假设，就是两个孩子都已经是大顶堆
package demo

import "fmt"

func MaxHeapFixDown(list []int, i int, n int) {
	var (
		parent int
		kid    int
	)
	parent = i
	// (i+1)*2 和 (i+1)*2 -1 ：为孩子节点
	kid = (i+1)*2 - 1
	for kid < n {
		// 取孩子最大的一个
		if kid+1 < n && list[kid+1] > list[kid] {
			kid++
		}
		// 最大的孩子和父亲比较
		if list[parent] < list[kid] {
			// 如果孩子大，则交换
			list[parent], list[kid] = list[kid], list[parent]
			// 父亲,孩子换人
			parent = kid
			kid = (parent+1)*2 - 1
			continue
		} else {
			// 如果父亲大，则退出，主要是假设在两个子树本来都是大顶堆
			return
		}
	}
}

// 大顶堆生成
func CreateMaxHeap(list []int) {
	// 从最小的非叶子节点开始生成。
	// 这样就会一直处在 两个孩子都已经是大顶堆 的前提
	// n := (len(list) - 1) / 2
	for n := (len(list) - 1) / 2; n >= 0; n-- {
		MaxHeapFixDown(list, n, len(list))
	}
}

// 从小到大排序
func HeapSort(list []int) {

	//1.大顶堆生成
	CreateMaxHeap(list)
	// 根与最后一个交换,此时我们概念中的堆的大小将减一
	// 也就是少一个节点不在大顶堆的范围内
	for n := len(list) - 1; n > 0; n-- {
		list[0], list[n] = list[n], list[0]
		// 因为动了根节点，对根节点进行处理
		MaxHeapFixDown(list, 0, n-1)
	}

}
func testHeapSort() {
	var inputList = []int{5, 3, 8, 22, 76, 1, 31, 55}
	// var inputList = []int{}
	// inputList = CreateList(inputList, 10)

	HeapSort(inputList)
	fmt.Println(inputList)
}
