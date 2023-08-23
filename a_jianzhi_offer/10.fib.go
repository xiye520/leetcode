package jianzhi_offer

/**
剑指 Offer 10- I. 斐波那契数列
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

示例 1：
输入：n = 2
输出：1

示例 2：
输入：n = 5
输出：5

https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/?envType=study-plan-v2&envId=coding-interviews
*/

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	for i := 2; i <= n; i++ {
		ans[i] = (ans[i-1] + ans[i-2]) % 1000000007
	}

	return ans[n]
}

func fib3(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
