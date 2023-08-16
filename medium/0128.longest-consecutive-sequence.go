package medium

/*
* 128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-interview-150
*/
func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	maxLength := 1
	for _, v := range nums {
		// 找到连续数列的起始位置
		if !m[v-1] {
			count := 1
			for m[v+1] {
				count++
				v++
			}
			maxLength = max(maxLength, count)
		}
	}

	return maxLength
}
