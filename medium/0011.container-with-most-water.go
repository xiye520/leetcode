package medium

import "github.com/xiye520/leetcode/util"

/*
11.盛最多水的容器

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
(i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
*/
func maxArea(height []int) (result int) {
	for i, j := 0, len(height)-1; i < j; {
		now := (j - i) * util.Min(height[i], height[j])
		result = util.Max(result, now)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}
