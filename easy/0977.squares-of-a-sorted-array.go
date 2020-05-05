package easy

import (
	"sort"
)

/*
给定一个按非递减顺序排序的整数数组 A，返回每个数字的平方组成的新数组，要求也按非递减顺序排序。
示例 1：

输入：[-4,-1,0,3,10]
输出：[0,1,9,16,100]
示例 2：

输入：[-7,-3,2,3,11]
输出：[4,9,9,49,121]

链接：https://leetcode-cn.com/problems/squares-of-a-sorted-array
*/
// 直接算完，然后用排序，时间复杂度0（n*lgn）
func sortedSquares(A []int) []int {
	var ret = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		ret[i] = A[i] * A[i]
	}

	sort.Ints(ret)
	return ret
}

// 因为原数组是有序的，所以只需要使用双指针前后标记，移位计算，即可做到O(n)时间复杂度
func sortedSquares2(A []int) []int {
	var ret = make([]int, len(A))
	start, end := 0, len(A)-1
	index := len(A) - 1
	for start <= end {
		if A[start]*A[start] < A[end]*A[end] {
			ret[index] = A[end] * A[end]
			end--
		} else {
			ret[index] = A[start] * A[start]
			start++
		}
		index--
	}

	return ret
}
