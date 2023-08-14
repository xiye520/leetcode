package easy

import (
	"fmt"
	"strconv"
)

/** 228. 汇总区间
给定一个  无重复元素 的 有序 整数数组 nums 。
返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
列表中的每个区间范围 [a,b] 应该按如下格式输出：
"a->b" ，如果 a != b
"a" ，如果 a == b

示例 1：
输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"

示例 2：
输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

*/
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	l := len(nums)
	if l == 0 {
		return res
	}
	if len(nums) == 1 {
		return append(res, strconv.Itoa(nums[0]))
	}
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		m[v] = true
	}
	start := nums[0]
	needAppend := false
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]-1] {
			start = nums[i]
		}
		// 有连续数据
		if m[nums[i]+1] {
			needAppend = true
			continue
		} else {
			if needAppend {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i]))
			} else {
				res = append(res, strconv.Itoa(start))
			}
			needAppend = false
		}
	}

	return res
}
