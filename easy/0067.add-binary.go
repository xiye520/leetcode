package easy

import "strconv"

/***
67. 二进制求和
给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。

示例 1：
输入:a = "11", b = "1"
输出："100"

示例 2：
输入：a = "1010", b = "1011"
输出："10101"

https://leetcode.cn/problems/add-binary/description/?envType=study-plan-v2&envId=top-interview-150
*/

func addBinary(a string, b string) string {
	alen, blen := len(a), len(b)
	n := alen
	if alen < blen {
		n = blen
	}
	ans := ""
	carry := 0
	for i := 0; i < n; i++ {
		if i < alen {
			carry += int(a[alen-i-1]) - '0'
		}
		if i < blen {
			carry += int(b[blen-i-1]) - '0'
		}
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
