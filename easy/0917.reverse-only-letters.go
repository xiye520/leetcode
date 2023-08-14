package easy

import "unicode"

/**
  给你一个字符串 s ，根据下述规则反转字符串：

  所有非英文字母保留在原有位置。
  所有英文字母（小写或大写）位置反转。
  返回反转后的 s 。

  示例 1：
  输入：s = "ab-cd"
  输出："dc-ba"

  示例 2：
  输入：s = "a-bC-dEf-ghIj"
  输出："j-Ih-gfE-dCba"

  示例 3：
  输入：s = "Test1ng-Leet=code-Q!"
  输出："Qedo1ct-eeLg=ntse-T!"

  来源：力扣（LeetCode）
  链接：https://leetcode.cn/problems/reverse-only-letters
*/
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		// 判断左侧是否字符串
		if !unicode.IsLetter(rune(s[left])) {
			left++
			continue
		}
		// 判断右侧是否字符串
		if !unicode.IsLetter(rune(s[right])) {
			right--
			continue
		}
		ans[left], ans[right] = ans[right], ans[left]
		left++
		right--
	}

	return string(ans)
}
