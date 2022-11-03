package hard

import (
	"github.com/xiye520/leetcode/data/kit"
)

/*
32. 最长有效括号
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

/*
1、先把 -1 放入栈内。（至于为什么？看到后面你就知道了）
2、、对于遇到的每个 '(' ，我们将它的下标放入栈中。
3、对于遇到的每个 ‘)’ ，我们弹出栈顶的元素并将当前元素的下标与弹出元素下标作差，得出当前有效括号字符串的长度。
通过这种方法，我们继续计算有效子字符串的长度，并最终返回最长有效子字符串的长度。
*/
// 2.使用栈来做
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
				n = max(n, i-stack.Peek()) // 用当前下标减去当前栈顶的元素
			}
		}
	}

	return n
}

/*
我们在从从左到右遍历字符串的过程中，用 left 记录 '(' 的数量，用 right 记录 ')' 的数量。并且在遍历的过程中：
1、如果 left == right，显然这个时候 right 个 ')' 都将一定能够得到匹配。所以当前的有效括号长度为 2 * right。然后更新 max。
2、如果 left < right，显然这个时候部分 ')' 一定得不到匹配，此时我们把 left 和 right 都置为 0。
*/
// 3.时间复杂度为 O(n)，空间复杂度为 O(1)
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
