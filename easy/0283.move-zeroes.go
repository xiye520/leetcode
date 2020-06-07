package easy

/*
283. 移动零

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*/
func moveZeroes(nums []int) {
	count := 0 // 计算零的个数
	for i := range nums {
		if nums[i] == 0 {
			count++
		}
	}

	start := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[start] = nums[i]
			start++
		}
	}
	// 末尾填零
	for i := 0; i < count; i++ {
		nums[len(nums)-1-i] = 0
	}
}

func moveZeroes2(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 找到数组中第一个0位置
	p := 0
	for p < n && nums[p] != 0 {
		p++
	}
	//
	for i := p + 1; i < n; i++ {
		if nums[i] != 0 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
}
