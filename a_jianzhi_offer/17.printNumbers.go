package jianzhi_offer

/**
剑指 Offer 17. 打印从1到最大的n位数
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]

说明：
用返回一个整数列表来代替打印
n 为正整数

https://leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/description/?envType=study-plan-v2&envId=coding-interviews
*/

func printNumbers(n int) []int {
	max := pow(n) - 1
	ans := make([]int, max)
	for i := 1; i <= max; i++ {
		ans[i-1] = i
	}
	return ans
}

func pow(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= 10
	}
	if res == 1 {
		return 0
	}
	return res
}
