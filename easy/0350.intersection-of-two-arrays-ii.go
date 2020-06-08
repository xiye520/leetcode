package easy

/*
350. 两个数组的交集 II

给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, 0)
	m2 := make(map[int]int, 0)
	for _, v := range nums1 {
		m1[v]++
	}

	for _, v := range nums2 {
		m2[v]++
	}

	ret := make([]int, 0, len(m1)+len(m2))
	for k, v := range m1 {
		if v2, ok := m2[k]; ok {
			c := min(v, v2)
			for i := 0; i < c; i++ {
				ret = append(ret, k)
			}
		}
	}

	return ret
}

func intersect2(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	// m1 是较短的字典，会快一些
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res
}

// 清理重复的值，也便于查询
func getInts(a []int) map[int]int {
	res := make(map[int]int, len(a))

	for i := range a {
		res[a[i]]++
	}

	return res
}
