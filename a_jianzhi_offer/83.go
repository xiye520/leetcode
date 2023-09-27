package jianzhi_offer

func permute(nums []int) [][]int {
	l := len(nums)
	if l == 1 {
		return [][]int{{nums[0]}}
	}
	if l == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	var res [][]int
	for i := range nums {
		nums[0], nums[i] = nums[i], nums[0]
		ce := permute(nums[1:])
		for _, arr := range ce {
			r := make([]int, 0, l)
			r = append(r, nums[0])
			r = append(r, arr...)
			res = append(res, r)
		}

		nums[0], nums[i] = nums[i], nums[0]
	}

	return res
}
