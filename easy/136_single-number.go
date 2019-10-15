package easy

/*
136. 只出现一次的数字

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4
*/
func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}

func singleNumber2(nums []int) int {
	m := make(map[int]struct{}, 0)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
			continue
		}
		m[n] = struct{}{}
	}

	for k, _ := range m {
		return k
	}

	return 0
}

func singleNumber3(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		if m[v] == true {
			m[v] = false
		} else {
			m[v] = true
		}
	}
	//无重复值的数,m中的值为true
	var k int
	for k, v := range m {
		if v == true {
			return k
		}
	}
	return k
}
