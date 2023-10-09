package easy

import "sort"

/** 2578. 最小和分割
给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：

num1 和 num2 直接连起来，得到 num 各数位的一个排列。
换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
num1 和 num2 可以包含前导 0 。
请你返回 num1 和 num2 可以得到的和的 最小 值。

注意：
num 保证没有前导 0 。
num1 和 num2 中数位顺序可以与 num 中数位顺序不同。


示例 1：
输入：num = 4325
输出：59
解释：我们可以将 4325 分割成 num1 = 24 和 num2 = 35 ，和为 59 ，59 是最小和。

示例 2：
输入：num = 687
输出：75
解释：我们可以将 687 分割成 num1 = 68 和 num2 = 7 ，和为最优值 75 。

*/

func splitNum(num int) int {
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}
	sort.Ints(arr)

	nums1 := 0
	nums2 := 0
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			nums1 = nums1*10 + arr[i]
		} else {
			nums2 = nums2*10 + arr[i]
		}
	}

	return nums1 + nums2
}

func splitNum2(num int) int {
	cnt := [10]int{}
	n := 0
	for ; num > 0; num /= 10 {
		cnt[num%10]++
		n++
	}
	ans := [2]int{}
	for i, j := 0, 0; i < n; i++ {
		for cnt[j] == 0 {
			j++
		}
		cnt[j]--
		ans[i&1] = ans[i&1]*10 + j
	}
	return ans[0] + ans[1]
}
