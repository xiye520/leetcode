package easy

import "math"

/** 643. 子数组最大平均数 I
请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
任何误差小于 10-5 的答案都将被视为正确答案。

示例 1：
输入：nums = [1,12,-5,-6,50,3], k = 4
输出：12.75
解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75

示例 2：
输入：nums = [5], k = 1
输出：5.00000

https://leetcode.cn/problems/maximum-average-subarray-i/
*/
func findMaxAverage(nums []int, k int) float64 {
	ans := float64(math.MinInt32)
	sum := 0.00
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	ans = math.Max(ans, sum/float64(k))
	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		ans = math.Max(ans, sum/float64(k))
	}

	return ans
}
