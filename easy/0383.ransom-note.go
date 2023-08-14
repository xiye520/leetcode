package easy

/*
* 383. 赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：
输入：ransomNote = "a", magazine = "b"
输出：false

示例 2：
输入：ransomNote = "aa", magazine = "ab"
输出：false

示例 3：
输入：ransomNote = "aa", magazine = "aab"
输出：true
*/
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[uint8]int, len(ransomNote))
	n := make(map[uint8]int, len(magazine))
	for i := range ransomNote {
		m[ransomNote[i]]++
	}
	for i := range magazine {
		n[magazine[i]]++
	}
	for k, v := range m {
		if n[k] < v {
			return false
		}
	}

	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := make(map[uint8]int, 26)

	for i := range magazine {
		m[magazine[i]]++
	}
	for i := range ransomNote {
		m[ransomNote[i]]--
		if m[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
