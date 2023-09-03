package medium

/** 1004. 最大连续1的个数 III
给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。

示例 1：
输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。

示例 2：
输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。

https://leetcode.cn/problems/max-consecutive-ones-iii/
*/
func longestOnes(nums []int, k int) int {
	ans := 0
	m := make(map[int]int, len(nums))
	start := 0
	for end := range nums {
		m[nums[end]]++
		if m[0] <= k {
			ans = max(ans, end-start+1)
		}
		for m[0] > k {
			m[nums[start]]--
			start++
		}
	}

	return ans
}

func longestOnes2(nums []int, k int) int {
	ans := 0
	m := [2]int{}
	start := 0
	for end := range nums {
		m[nums[end]]++
		if m[0] <= k {
			ans = max(ans, end-start+1)
		}
		for m[0] > k {
			m[nums[start]]--
			start++
		}
	}

	return ans
}
