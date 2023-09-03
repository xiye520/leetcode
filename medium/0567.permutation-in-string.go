package medium

/*** 567. 字符串的排列
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。

示例 1：
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").

示例 2：
输入：s1= "ab" s2 = "eidboaoo"
输出：false

https://leetcode.cn/problems/permutation-in-string/
*/
func checkInclusion(s1 string, s2 string) bool {
	s1l, s2l := len(s1), len(s2)
	if s1l > s2l {
		return false
	}
	s1count, s2count := [26]int{}, [26]int{}
	for i := range s1 {
		s1count[s1[i]-'a']++
		s2count[s2[i]-'a']++
	}
	if s1count == s2count {
		return true
	}
	for i := 0; i < s2l-s1l; i++ {
		s2count[s2[i]-'a']--
		s2count[s2[i+s1l]-'a']++
		if s1count == s2count {
			return true
		}
	}

	return false
}
