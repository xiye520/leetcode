package medium

/** 46. 全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]
https://leetcode.cn/problems/permutations/description/
*/

var (
	res  [][]int
	path []int
	st   []bool // state的缩写
)

func permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))
	dfspermute(nums, 0)
	return res
}

func dfspermute(nums []int, cur int) {
	if cur == len(nums) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if !st[i] {
			path = append(path, nums[i])
			st[i] = true
			dfspermute(nums, cur+1)
			st[i] = false
			path = path[:len(path)-1]
		}
	}
}

/*
遍历数组里的每个元素，与第一个元素交换，表示以它开头，然后递归执行第一个元素后面的子数组，
得到的二维数组就是它的子问题，所以遍历得到的二维数组，加到这个元素的后面就是一个全排列；
注意在处理完一个元素后要交换回来
*/

func permute3(nums []int) [][]int {
	l := len(nums)
	if l == 1 {
		return [][]int{{nums[0]}}
	}
	if l == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	var res [][]int
	for i := range nums {
		// 遍历数组元素，以某个元素为起点
		nums[0], nums[i] = nums[i], nums[0]
		ce := permute3(nums[1:])
		for _, arr := range ce {
			var r []int
			r = append(r, nums[0])
			r = append(r, arr...)
			res = append(res, r)
		}
		// 要换回来
		nums[0], nums[i] = nums[i], nums[0]
	}

	return res
}
