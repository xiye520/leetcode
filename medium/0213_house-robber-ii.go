package medium

import "github.com/xiye520/leetcode/util"

/*
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [2,3,2]
输出: 3
解释: 你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。
示例 2:

输入: [1,2,3,1]
输出: 4
解释: 你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
     偷窃到的最高金额 = 1 + 3 = 4 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 环状排列意味着第一个房子和最后一个房子中只能选择一个偷窃，因此可以把此环状排列房间问题约化为两个单排排列房间子问题：
// 1、在不偷窃第一个房子的情况下（即 nums[1:]），最大金额是p1；
// 2、在不偷窃最后一个房子的情况下（即 nums[:n-1]），最大金额是p2;
// 综合偷窃最大金额： 为以上两种情况的较大值，即 util.Max(p1,p2) 。
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}
	return util.Max(tool(nums[:len(nums)-1]), tool(nums[1:]))
}

func tool(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return util.Max(nums[0], nums[1])
	}

	f := make([]int, len(nums))
	f[0] = nums[0]
	f[1] = util.Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		f[i] = util.Max(f[i-2]+nums[i], f[i-1])
	}

	return f[len(nums)-1]
}
