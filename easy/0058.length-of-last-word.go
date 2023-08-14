package easy

import "strings"

/** 58. 最后一个单词的长度
给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。



示例 1：
输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为5。

示例 2：
输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为4。

示例 3：
输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为6的“joyboy
https://leetcode.cn/problems/length-of-last-word/
*/

func lengthOfLastWord(s string) int {
	s = strings.ToLower(s)
	count := 0
	start := true
	for i := len(s) - 1; i >= 0; i-- {
		if isCharacter(s[i]) {
			start = false
			count++
			continue
		}
		if !start {
			break
		}
	}
	return count
}

func isCharacter(c uint8) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}
	return false
}
