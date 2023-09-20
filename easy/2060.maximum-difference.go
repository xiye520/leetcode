package easy

func maximumDifference(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	res := -1
	preMin := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > preMin {
			res = max(res, nums[i]-preMin)
		} else {
			preMin = nums[i]
		}
	}

	return res
}
