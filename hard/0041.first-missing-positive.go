package hard

import "math"

/*
41. 缺失的第一个正数
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1

提示：
你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
https://leetcode-cn.com/problems/first-missing-positive/
*/
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func firstMissingPositive2(nums []int) int {
	// 如果只有1就返回2
	if len(nums) == 1 {
		if nums[0] == 1 {
			return 2
		} else {
			return 1
		}
	}

	// 如果没有1返回1成功
	var flag_one int
	for _, v := range nums {
		if v == 1 {
			flag_one = 1
		}
	}
	if flag_one == 0 {
		return 1
	}

	// 把所有的负数和大于n的数字置1
	for k, v := range nums {
		if v <= 0 {
			nums[k] = 1
		}
		if v > len(nums) {
			nums[k] = 1
		}
	}

	// 把所有的出现数字中对应的value值乘以-1，代表这个数字出现了
	for _, v := range nums {
		if int(math.Abs(float64(v))) == len(nums) {
			nums[0] = int(math.Abs(float64(v)))
		} else if nums[int(math.Abs(float64(v)))] > 0 {
			nums[int(math.Abs(float64(v)))] *= -1
			if nums[0] < int(math.Abs(float64(v))) {
				nums[0] = int(math.Abs(float64(v)))
			}
		}
	}

	// 返回结果
	for i := 1; i < len(nums); i++ {
		if nums[i] > 0 {
			return i
		}
	}
	return nums[0] + 1
}

func firstMissingPositive3(nums []int) int {
	l := len(nums)
	contains := 0
	for i := 0; i < l; i++ { //检查是否包含1
		if nums[i] == 1 {
			contains++
			break
		}
	}

	if contains == 0 { //不包含1即答案
		return 1
	}
	if l == 1 { //包含且长度为1，  2即答案
		return 2
	}

	for i := 0; i < l; i++ { //把所有大数,负数，全部转换为1，因为值必定 res <= l+1
		if nums[i] <= 0 || l < nums[i] {
			nums[i] = 1
		}
	}

	for i := 0; i < l; i++ { //以符号的正负 当作hash表
		index := int(math.Abs(float64(nums[i]))) //获取需要改变的索引
		if index == l {
			nums[0] = -int(math.Abs(float64(nums[0])))
		} else {
			nums[index] = -int(math.Abs(float64(nums[index])))
		}
	}

	for i := 1; i < l; i++ { //不存在 即答案
		if nums[i] > 0 {
			return i
		}
	}

	if nums[0] > 0 { // 不存在 即答案
		return l
	}
	return l + 1 //都存在，后序l+1为答案
}
