package main

import "fmt"

// 1.闭包实现
func fabonacci_pack() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

// 2.递归实现
func fabo_recursion(n int) int {
	if n <= 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fabo_recursion(n-1) + fabo_recursion(n-2)
}

// 3.channel实现
func fabo_channel(output chan<- int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case output <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("stop fabonacci ")
			return
		}
	}
}

// 4.动态规划
func fabp_dp(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
