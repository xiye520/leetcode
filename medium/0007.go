package medium

import "math"

/*
* 7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。

示例 1：
输入：x = 123
输出：321

示例 2：
输入：x = -123
输出：-321

示例 3：
输入：x = 120
输出：21

示例 4：ß
输入：x = 0
输出：0

https://leetcode.cn/problems/reverse-integer/description/
*/
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	res := 0
	for x > 0 {
		res = res * 10
		res += x % 10
		x = x / 10
	}
	res = res * sign
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}

func reverse2(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	// 还原 x 的符号到 res
	res = sign * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}
