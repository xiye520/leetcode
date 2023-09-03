package jianzhi_offer

/** 剑指 Offer 42. 连续子数组的最大和
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/description/
*/
func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
