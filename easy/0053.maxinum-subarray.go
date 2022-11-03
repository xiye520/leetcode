package easy

import (
	"github.com/xiye520/leetcode/util"
	"math"
)

/*
53. 最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

链接：https://leetcode-cn.com/problems/maximum-subarray
*/
func maxSubArray(nums []int) int {
	// dp
	// 1.状态定义 F(i)表示下标为i时连续数组的最大和
	// 2.状态转移方程 F(i) = util.Max(F(i-1)+nums[i],nums[i]);
	// 3.初始状态 F(0) = nums[0];
	// 4.返回值 F(nums.size()-1)
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	fn := nums[0]
	for i := 1; i < len(nums); i++ {
		fn = util.Max(nums[i], fn+nums[i])
		res = util.Max(fn, res)
	}

	return res
}

func maxSubArray2(nums []int) int {

	maxSum := 0
	finalMaxSum := math.MinInt32
	for _, num := range nums {
		maxSum += num

		if maxSum < num {
			maxSum = num
		}

		if maxSum > finalMaxSum {
			finalMaxSum = maxSum
		}
	}

	return finalMaxSum
}
