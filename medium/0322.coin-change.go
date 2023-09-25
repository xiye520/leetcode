package medium

/*
* 322. 零钱兑换
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。

示例 1：
输入：coins = [1, 2, 5], amount = 11
输出：3
解释：11 = 5 + 5 + 1

示例 2：
输入：coins = [2], amount = 3
输出：-1

示例 3：
输入：coins = [1], amount = 0
输出：0

https://leetcode.cn/problems/coin-change/
*/
func coinChange(coins []int, amount int) int {
	// dp[i] 更换金额为 i 时，最小的零钱数量
	// res = dp[amount]
	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i] > dp[i-c]+1 {
				dp[i] = dp[i-c] + 1
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

//func coinChange3(coins []int, amount int) int {
//	max := amount + 1
//	dp := make([]int, amount+1)
//	for i := range dp {
//		dp[i] = max
//	}
//	dp[0] = 0
//
//	sort.Ints(coins)
//
//	for i := 1; i <= amount; i++ {
//		for _, c := range coins {
//			if i-c < 0 {
//				break
//			}
//			dp[i] = min(dp[i], dp[i-c]+1)
//		}
//	}
//
//	if dp[amount] > amount {
//		return -1
//	}
//
//	return dp[amount]
//}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i {
				if dp[i] > dp[i-c]+1 {
					dp[i] = dp[i-c] + 1
				}
			}

		}

	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
