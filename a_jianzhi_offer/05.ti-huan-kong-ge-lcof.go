package jianzhi_offer

import "strings"

/*
*
剑指 Offer 05. 替换空格
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

https://leetcode.cn/problems/ti-huan-kong-ge-lcof/?envType=study-plan-v2&envId=coding-interviews
*/
func replaceSpace(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
