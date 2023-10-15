package medium

/*
* 137. 只出现一次的数字 II
给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。

示例 1：
输入：nums = [2,2,3,2]
输出：3

示例 2：
输入：nums = [0,1,0,1,0,1,99]
输出：99

https://leetcode.cn/problems/single-number-ii/
*/
func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums)/3)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v <= 1 {
			return k
		}
	}
	return -1
}

func singleNumber2(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}
