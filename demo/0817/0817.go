package main

import (
	"fmt"
)

/*
题目：
给一个日期：20200202，观察这个日期，同时满足两个特征：
1.左右对称
2.包含数字0，除数字0外，只有一个非零数字

编码实现以下两项：
1.输入任意日期，判断是否满足以上条件，测试用例：20200202，21211212
2.输出自19700101至今所有符合条件的日期
*/
func getDates(start int) []string {

	return GetBetweenDates("19700101", "20230817")
}

func getDays(syear, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
	// 处理闰年
	default:
		panic(fmt.Sprintf("输入年份不规范：%d", month))

	}
	return 0
}

func isValid(s string) bool {
	if !isMatch(s) {
		return false
	}
	return isOnce(s)
}

// 判断是否左右对称
func isMatch(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// 包含数字0，除数字0外，只有一个非零数字
func isOnce(s string) bool {
	m := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		m[s[i]] = true
		if s[i] == '0' {
			continue
		}

		if len(m) > 2 {
			return false
		}
	}
	if !m['0'] {
		return false
	}

	return true
}

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数
测试用例：
nums = [1,2,3,4,5,6,7], k = 3，输出[5，6，7，1，2，3，4]
nums = [1,2,3], k = 4，输出[3,1,2]
*/
func transfer(nums []int, k int) []int {
	temp := make([]int, len(nums))
	k = k % len(nums)

	for i, _ := range nums {
		if i+k < len(nums) {
			temp[i+k] = nums[i]
		} else {
			temp[i+k-len(nums)] = nums[i]
		}
	}
	copy(nums, temp)

	return nums
}
