package easy

/*
* 205. 同构字符串
给定两个字符串 s 和 t ，判断它们是否是同构的。

如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

示例 1:
输入：s = "egg", t = "add"
输出：true

示例 2：
输入：s = "foo", t = "bar"
输出：false

示例 3：
输入：s = "paper", t = "title"
输出：true

示例 2：
输入：s = "bacd", t = "baba"
输出：false
https://leetcode.cn/problems/isomorphic-strings/
*/
func isIsomorphic(s string, t string) bool {
	return compare(s, t) && compare(t, s)
}

func compare(s string, t string) bool {
	m := make(map[uint8]uint8, len(s))
	for i, _ := range s {
		if c, ok := m[s[i]]; ok {
			if c != t[i] {
				return false
			}
		} else {
			m[s[i]] = t[i]
		}
	}

	return true
}

func isIsomorphic2(s, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
