package easy

/*
* 697. 数组的度
给定一个非空且只包含非负数的整数数组 nums，数组的 度 的定义是指数组里任一元素出现频数的最大值。
你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

示例 1：
输入：nums = [1,2,2,3,1]
输出：2
解释：
输入数组的度是 2 ，因为元素 1 和 2 的出现频数最大，均为 2 。
连续子数组里面拥有相同度的有如下所示：
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组 [2, 2] 的长度为 2 ，所以返回 2 。

示例 2：
输入：nums = [1,2,2,3,1,4,2]
输出：6
解释：
数组的度是 3 ，因为元素 2 重复出现 3 次。
所以 [2,2,3,1,4,2] 是最短子数组，因此返回 6 。
https://leetcode.cn/problems/degree-of-an-array/description/
*/
type entry struct{ cnt, l, r int }

func findShortestSubArray(nums []int) (ans int) {
	m := map[int]entry{}
	for i, v := range nums {
		if e, has := m[v]; has {
			// 数据已存在，增加计数，更新终点（右）坐标
			e.cnt++
			e.r = i
			m[v] = e
		} else {
			// 数据不存在，初始化计数，起点、终点坐标
			m[v] = entry{1, i, i}
		}
	}
	maxCnt := 0
	for _, e := range m {
		if e.cnt > maxCnt {
			maxCnt, ans = e.cnt, e.r-e.l+1
		} else if e.cnt == maxCnt {
			ans = min(ans, e.r-e.l+1)
		}
	}
	return
}
