package medium

import (
	"bytes"
	"strconv"
	"unicode"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

链接：https://leetcode-cn.com/problems/decode-string
*/
type Pair struct {
	nextNum int
	thisStr string
}

func multiplyStr(n int, s string) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(s)
	}
	return buffer.String()
}

func decodeString(s string) string {
	num := 0
	curStr := ""
	stack := []Pair{}

	for _, c := range s {
		if unicode.IsDigit(c) { // end of string
			n, _ := strconv.Atoi(string(c))
			num = num*10 + n
		} else if c == '[' { // end of number
			stack = append(stack, Pair{num, curStr})
			num, curStr = 0, ""
		} else if c == ']' {
			pair := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curStr = pair.thisStr + multiplyStr(pair.nextNum, curStr)
		} else {
			curStr = curStr + string(c)
		}
	}

	return curStr
}

func decodeString2(s string) string {
	// 2[abc]3[cd]ef
	num := 0
	cur := ""
	stack := []Pair{}
	for _, c := range s {
		switch {
		case unicode.IsDigit(c): //
			n, _ := strconv.Atoi(string(c))
			num += num*10 + n

		case c == '[': // 前面的数字到此结束
			stack = append(stack, Pair{nextNum: num, thisStr: cur})
			num = 0
			cur = ""
		case c == ']':
			last := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			cur = last.thisStr + multiplyStr(last.nextNum, cur) // 将此前的拼接上来
		default:
			cur = cur + string(c)
		}
	}

	return cur
}
