package easy

import "fmt"

/*
169.求众数

给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:

输入: [3,2,3]
输出: 3

示例 2:

输入: [2,2,1,1,1,2,2]
输出: 2

*/

func majorityElement(nums []int) int {
	count := make(map[int]int, 0)
	for _, n := range nums {
		fmt.Println(n)
		fmt.Println(count)
		count[n]++
		fmt.Println(count)
	}

	preTimes := 0
	result := 0
	for num, times := range count {
		if preTimes < times {
			preTimes = times
			result = num
		}
	}

	return result
}

func majorityElement2(nums []int) int {
	// 根据题意 len[nums] > 0 且 出现次数大于 n/2 的元素存在。
	x, t := nums[0], 1

	for i := 1; i < len(nums); i++ {
		switch {
		case x == nums[i]:
			t++
		case t > 0:
			t--
		default:
			// 此时 x != nums[i] 且 t == 0
			// 可知 i 必定为 偶数
			// 假设 nums 中出现最多的元素是 z，其出现次数为 zn > n/2
			// 在 nums[:i] 中，z 出现次数 <=i/2
			// 那么，在 nums[i:] 中，z 出现的次数 >= zn - i/2 > n/2 - i/2 = (n-i)/2
			// 即 z 在 nums[i:] 中出现的次数，依然超过了 len(nums[i:])/2
			x = nums[i]
			t = 1
		}
	}

	return x
}

func majorityElement3(nums []int) int {
	half := len(nums) / 2
	count := make(map[int]int, 0)
	for _, n := range nums {
		count[n]++
		if count[n] > half {
			return n
		}
	}

	return 0
}
