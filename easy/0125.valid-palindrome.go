package easy

import (
	"bytes"
	"fmt"
	"strings"
)

/*
125. 验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

ASCII码：
0-48
9-57

a-97
z-122
*/

func isPalindrome_(s string) bool {
	//s := "A man, a plan, a canal: Panama"
	s = strings.ToLower(s)
	bf := bytes.Buffer{}
	for _, v := range s {
		if (v >= 48 && v <= 57) || (v >= 97 && v <= 122) {
			bf.WriteString(fmt.Sprintf("%c", v))
		}
	}
	s = bf.String()

	if len(s) < 2 {
		return true
	}
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--

	}

	return true
}

func isPalindrome_2(s string) bool {
	s = strings.ToLower(s)
	fmt.Println(s)
	start, end := 0, len(s)-1
	for start < end {
		for start < end && !isChar(s[start]) {
			start++
		}

		for start < end && !isChar(s[end]) {
			end--
		}

		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}

func isPalindrome_3(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// 判断 c 是否是需要检查的字符
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
