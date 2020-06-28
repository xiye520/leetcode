package medium

func maxProduct(nums []int) int {
	for len(nums) == 0 {
		return 0
	}

	for len(nums) == 1 {
		return nums[0]
	}
	mi, ma, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			mi, ma = ma, mi
		}
		ma = max(ma*nums[i], nums[i])
		mi = min(mi*nums[i], nums[i])
		res = max(res, ma)
	}

	return res
}
