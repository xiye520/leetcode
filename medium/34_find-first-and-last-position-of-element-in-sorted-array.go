package medium

import "fmt"

/*
34. 在排序数组中查找元素的第一个和最后一个位置

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func searchRange(nums []int, target int) []int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		fmt.Println("mid", mid, "value", nums[mid])
		if nums[mid] == target {
			//顺着这个mid前后找
			fmt.Println("mid", mid, "value", nums[mid])
			var start, end int
			for start = mid; start >= low; start-- {
				if nums[start] != target {
					break
				}
			}

			for end = mid; end <= high; end++ {
				if nums[end] != target {
					break
				}
			}
			return []int{start + 1, end - 1}
		} else if nums[mid] < target { // 中值小于目标值
			low = mid + 1
		} else { // 中值大于目标值
			high = mid - 1
		}
	}

	return []int{-1, -1}
}

func searchRange2(nums []int, target int) []int {
	index := search(nums, target)
	if index == -1 {
		return []int{-1, -1}
	}

	first := index
	for {
		f := search(nums[:first], target)
		if f == -1 {
			break
		}
		first = f
	}

	last := index
	for {
		l := search(nums[last+1:], target)
		if l == -1 {
			break
		}
		last = l + 1
	}

	return []int{first, last}
}

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
