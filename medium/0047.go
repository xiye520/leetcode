package medium

import "sort"

func permuteUnique(nums []int) [][]int {
	l := len(nums)
	if l == 1 {
		return [][]int{{nums[0]}}
	}
	if l == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	}
	var res [][]int
	m := make(map[int]bool, l)
	for i := range nums {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true
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

func permuteUnique2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	perm := []int{}
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}

// 在回溯法求解全排列的基础上，添加了一个 map 重复，回溯法经典的思想，
// 每次递归中维护一个 map 来检查是否在此阶段将两个相同的数交换了位置，若交换了则必定有所重复。
func permuteUnique3(nums []int) [][]int {
	res := [][]int{}
	dfspermuteUnique(nums, &res, 0)
	return res
}

func dfspermuteUnique(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		*res = append(*res, dump(nums))
	}

	m := map[int]int{}
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		dfspermuteUnique(nums, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		m[nums[i]] = 1
	}
}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
