package easy

/**485. 最大连续 1 的个数
给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。

示例 1：
输入：nums = [1,1,0,1,1,1]
输出：3
解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.

示例 2:
输入：nums = [1,0,1,1,0,1]
输出：2

https://leetcode.cn/problems/max-consecutive-ones/
*/
func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	start := 0
	count := 0
	for start < len(nums) {
		if nums[start] == 1 {
			count++
			ans = max(ans, count)
		} else {
			count = 0
		}
		start += 1
	}

	return ans
}
