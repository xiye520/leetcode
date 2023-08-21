package medium

/** 189. 轮转数组
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:
输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
*/

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}

	k = k % len(nums)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i+k < len(nums) {
			res[i+k] = nums[i]
		} else {
			res[i+k-len(nums)] = nums[i]
		}
	}
	copy(nums, res)
}
