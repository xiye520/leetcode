package easy

import "strings"

/**
290. 单词规律
给定一种规律 pattern 和一个字符串 s ，判断 s 是否遵循相同的规律。
这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 s 中的每个非空单词之间存在着双向连接的对应规律。

示例1:
输入: pattern = "abba", s = "dog cat cat dog"
输出: true

示例 2:
输入:pattern = "abba", s = "dog cat cat fish"
输出: false

示例 3:
输入: pattern = "aaaa", s = "dog cat cat dog"
输出: false

https://leetcode.cn/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150
*/

func wordPattern(pattern string, s string) bool {
	word2char := make(map[string]rune, len(pattern))
	char2word := make(map[rune]string, len(pattern))
	slice := strings.Fields(s)
	if len(pattern) != len(slice) {
		return false
	}
	for i, char := range pattern {
		if (word2char[slice[i]] > 0 && char2word[char] != slice[i]) || (char2word[char] != "" && word2char[slice[i]] != char) {
			return false
		}

		char2word[char] = slice[i]
		word2char[slice[i]] = char
	}

	return true
}
