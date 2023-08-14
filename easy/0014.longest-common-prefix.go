package easy

import "fmt"

/*
14. 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var count = 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			fmt.Println(j, i, len(strs[j]) < i)
			if len(strs[j]) <= i || strs[j][i] != strs[0][i] {
				return strs[0][:count]
			}
		}
		count++
	}

	return strs[0][:count]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	var prefix = strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

// 取出两个字符串中的最长公共前缀
func lcp(str1, str2 string) string {
	maxL := min(len(str1), len(str2))
	for i := 0; i < maxL; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}
	return str1[:maxL]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
