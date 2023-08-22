package jianzhi_offer

/*
*
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
示例：

输入：nums = [1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
*/
func exchange(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	start, end := 0, len(nums)-1
	for start < end {
		for start < len(nums)-1 && nums[start]%2 == 1 {
			start++
		}
		for end >= 0 && nums[end]%2 == 0 {
			end--
		}
		if start > end {
			break
		}
		nums[start], nums[end] = nums[end], nums[start]
	}

	return nums
}
