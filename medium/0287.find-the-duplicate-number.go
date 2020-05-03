package medium

/*
287. 寻找重复数
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:

输入: [1,3,4,2,2]
输出: 2

https://leetcode-cn.com/problems/find-the-duplicate-number/
*/

// https://leetcode-cn.com/problems/find-the-duplicate-number/solution/goyu-yan-fu-luo-yi-de-suan-fa-zhao-huan-ru-kou-by-/
// 弗洛伊德算法找环入口，时间复杂度O(n)，空间复杂度O(1)
func findDuplicate(nums []int) int {
	//1.初始化快慢指针
	slow, fast := 0, 0
	for {
		//2.慢指针走一步，快指针走两步
		slow = nums[slow]
		fast = nums[fast]
		fast = nums[fast]
		//3.快慢指针重合，完成弗洛伊得算法的第一步，找交点
		if slow == fast {
			break
		}
	}
	//4.判断入口点
	fast = 0
	for {
		//4.1 快指针从开始位置遍历，慢指针从交点位置遍历，相等的位置即为环入口
		fast = nums[fast]
		slow = nums[slow]
		if slow == fast {
			break
		}
	}

	return fast
}

// 使用map记录出现过的  时空复杂度均为O(n)
func findDuplicate2(nums []int) int {
	m := make(map[int]struct{}, len(nums)-1)
	m[nums[0]] = struct{}{}
	for i := 1; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = struct{}{}
	}

	return -1
}
