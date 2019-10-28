package easy

/*
905. 按奇偶排序数组

给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。
示例：

输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。


提示：

1 <= A.length <= 5000
0 <= A[i] <= 5000
*/

func sortArrayByParity(A []int) []int {
	start, end := 0, len(A)-1
	for {
		for ; start <= end; start++ {
			if A[start]%2 != 0 { // 奇数
				break
			}
		}

		for ; end >= start; end-- {
			if A[end]%2 == 0 { // 偶数
				break
			}
		}
		if start >= end {
			break
		}
		A[start], A[end] = A[end], A[start]
	}

	return A
}

func sortArrayByParity2(a []int) []int {
	i, j := 0, len(a)-1
	for {
		for i < j && a[i]%2 == 0 {
			i++
		}
		for i < j && a[j]%2 == 1 {
			j--
		}

		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			break
		}

	}
	return a
}
