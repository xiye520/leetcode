package easy

/*
217. 存在重复元素




题目描述
评论 (522)
题解(71)New
提交记录
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
在真实的面试中遇到过这道
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for i := range nums {
		if _, ok := m[nums[i]]; ok {
			return true
		}
		m[nums[i]] = 1
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}

	return false
}
