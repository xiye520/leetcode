package easy

import "math"

/*

387.字符串中的第一个唯一字符

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/

func firstUniqChar(s string) int {
	m := make(map[rune]int, 0)

	for _, v := range s {
		m[v] += 1
	}

	for i, v := range s {
		if m[v] == 1 {
			return i
		}
	}

	return -1
}

// 效率更高，比map高十倍的速度
func firstUniqChar2(s string) int {
	table := [26]int{}

	for i := 0; i < len(s); i++ {
		table[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if table[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar3(s string) int {
	cnt := [26]int{}
	for i, c := range s {
		k := c - 'a'
		if cnt[k] == 0 {
			cnt[k] = i + 1
		} else if cnt[k] > 0 {
			cnt[k] = math.MaxInt32
		}
	}
	min := math.MaxInt32
	for _, j := range cnt {
		if j > 0 && j < min {
			min = j
		}
	}
	if min == math.MaxInt32 {
		min = 0
	}
	return min - 1
}
