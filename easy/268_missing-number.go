package easy

/*
268. 缺失数字

给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
示例 1:
输入: [3,0,1]
输出: 2

示例 2:
输入: [9,6,4,2,3,5,7,0,1]
输出: 8
*/
func missingNumber(nums []int) int {
	// 应该是连续len(nums)+1个数，缺失了其中一个，故为总数-当前nums总数之差
	l := len(nums)
	real := (l + 1) * l / 2
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return real - sum
}

func missingNumber2(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		m[n] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}

	return 0
}
