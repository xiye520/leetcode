package medium

import (
	"sort"
)

/**
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1:
输入: [3,2,1,5,6,4], k = 2
输出: 5

示例 2:
输入: [3,2,3,1,2,4,5,5,6], k = 4
输出: 4

https://leetcode.cn/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=top-interview-150
*/

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func findKthLargest3(nums []int, k int) int {
	n := len(nums)
	return quickselect2(nums, 0, n-1, n-k)
}

func quickselect2(nums []int, left, right, k int) int {
	if left >= right {
		return nums[k]
	}
	pivot := partiton(nums, left, right)
	if pivot == k {
		return nums[k]
	}
	// 中间位置小于k，排序最右侧的数组即可
	if pivot < k {
		quickselect2(nums, pivot+1, right, k)
	} else {
		quickselect2(nums, left, pivot-1, k)
	}
	return nums[k]
}

func partiton(arr []int, left, right int) int {
	key := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < key {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}

	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}
