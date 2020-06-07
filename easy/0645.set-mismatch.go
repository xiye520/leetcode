package easy

/*
645. 错误的集合
集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。

给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

示例 1:

输入: nums = [1,2,2,4]
输出: [2,3]

https://leetcode-cn.com/problems/set-mismatch/
*/
func findErrorNums(nums []int) []int {
	fullSum := (len(nums) + 1) * len(nums) / 2
	m := make(map[int]struct{}, len(nums))
	var sum int
	var repeat int
	for _, v := range nums {
		if _, ok := m[v]; ok {
			repeat = v
		} else {
			m[v] = struct{}{}
		}
		sum += v
	}

	return []int{repeat, repeat + (fullSum - sum)}
}

// 根据传入数组的长度指定切片长度，用切片来完成set的功能，内存和速度都可以达到更优
func findErrorNums2(nums []int) []int {
	length := len(nums)
	templist := make([]int, length+1)

	res1, res2 := 0, 0

	for i := 0; i < length; i++ {
		templist[nums[i]]++
	}

	for i := 1; i < length+1; i++ {
		if templist[i] == 2 {
			res1 = i
		} else if templist[i] == 0 {
			res2 = i
		}
	}

	return []int{res1, res2}
}
