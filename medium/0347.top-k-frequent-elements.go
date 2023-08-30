package medium

/**
347. 前 K 个高频元素
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]

示例 2:
输入: nums = [1], k = 1
输出: [1]


提示：
1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的

进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
https://leetcode.cn/problems/top-k-frequent-elements/description/
*/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	maxCount := 0
	for _, v := range nums {
		m[v]++
		maxCount = max(maxCount, m[v])
	}
	bucket := make([][]int, maxCount+1)
	for k, count := range m {
		bucket[count] = append(bucket[count], k)
	}

	l := 0
	res := make([]int, 0, k)
	for i := maxCount; i > 0; i-- {
		for _, v := range bucket[i] {
			res = append(res, v)
			l++
			if l >= k {
				return res
			}
		}
	}

	return res
}
