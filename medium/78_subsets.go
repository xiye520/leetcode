package medium

/*
78. 子集

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	n := 1 << uint(len(nums))
	for i := 0; i < n; i++ {
		temp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if uint(i)>>uint(j)&1 == 1 {
				temp = append(temp, nums[j])
			}
		}
		result = append(result, temp)
	}

	return result
}

func subset2(nums []int) [][]int {
	result := make([][]int, 0)
	subsetsBT(&result, nums, []int{}, 0)
	return result
}
func subsetsBT(result *[][]int, nums []int, temp []int, start int) {
	//此处深拷贝temp，避免回溯的时候temp被修改后会影响之前保存的结果
	c := make([]int, len(temp))
	copy(c, temp)
	*result = append(*result, c)

	for i := start; i < len(nums); i++ {
		temp = append(temp, nums[i])
		subsetsBT(result, nums, temp, i+1) //不包含重复值
		temp = temp[:len(temp)-1]
	}
}
