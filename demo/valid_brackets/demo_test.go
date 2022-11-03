package main

import (
	"github.com/xiye520/leetcode/data/kit"
	"testing"
)

/*
给定一个只包括 '('，')'的字符串，判断字符串是否有效。注：空字符串属于有效字符串

示例 1:
输入: "(())"
输出: true

 实例 2：
 输入: "())("
输出: false
*/
func isValid(s string) bool {
	if s == "" {
		return true
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sum++
		} else {
			sum--
		}
		if sum < 0 {
			return false
		}
	}

	return sum == 0
}

func Test_isValid(t *testing.T) {
	par := map[string]bool{
		"(())": true,
		"())(": false,
		"":     true,
	}
	for k, v := range par {
		if isValid(k) != v {
			t.Fatal(k, v)
		}
	}
}

/*
三、最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

// 1.暴力解法
func longestValidParentheses(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		for now := 2; i+now <= len(s); now += 2 {
			if isValid(s[i : i+now]) {
				n = max(n, now)
			}
		}
	}
	return n
}

func Test_longestValidParentheses(t *testing.T) {

	//t.Log(longestValidParentheses("(()"))
	//t.Log(longestValidParentheses(")()())"))

	par := map[string]int{
		"(()":    2,
		")()())": 4,
		"":       0,
	}
	for k, v := range par {
		res := longestValidParentheses3(k)
		if res != v {
			t.Fatal(res, k, v)
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func longestValidParentheses2(s string) int {
	n := 0
	stack := kit.NewStack()
	stack.Push(-1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				n = max(n, i-stack.Peek())
			}
		}
	}

	return n
}

func longestValidParentheses3(s string) int {
	if len(s) < 2 {
		return 0
	}
	n := 0
	left, right := 0, 0
	// 从左到右
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			n = max(n, left*2)
		} else if left < right {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	// 从右到左
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			n = max(n, left*2)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return n
}
