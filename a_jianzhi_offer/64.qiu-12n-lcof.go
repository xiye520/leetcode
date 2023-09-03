package jianzhi_offer

/** 剑指 Offer 64. 求1+2+…+n
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

示例 1：
输入: n = 3
输出: 6

示例 2：
输入: n = 9
输出: 45

https://leetcode.cn/problems/qiu-12n-lcof/description/
*/
func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}
