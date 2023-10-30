package medium

import "sort"

/** 260. 只出现一次的数字 III
给你一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
你必须设计并实现线性时间复杂度的算法且仅使用常量额外空间来解决此问题。

示例 1：
输入：nums = [1,2,1,3,2,5]
输出：[3,5]
解释：[5, 3] 也是有效的答案。

示例 2：
输入：nums = [-1,0]
输出：[-1,0]

示例 3：
输入：nums = [0,1]
输出：[1,0]

https://leetcode.cn/problems/single-number-iii/?envType=daily-question&envId=2023-10-16
*/

func singleNumber260(nums []int) []int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}
	ans := make([]int, 0, 2)
	for k, v := range m {
		if v != 2 {
			ans = append(ans, k)
		}
		if len(ans) == 2 {
			sort.Ints(ans)
			return ans
		}
	}

	sort.Ints(ans)
	return ans
}
