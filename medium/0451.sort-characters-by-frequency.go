package medium

import (
	"bytes"
	"sort"
)

/*
* 451. 根据字符出现频率排序
给定一个字符串 s ，根据字符出现的 频率 对其进行 降序排序 。一个字符出现的 频率 是它出现在字符串中的次数。

返回 已排序的字符串 。如果有多个答案，返回其中任何一个。

示例 1:

输入: s = "tree"
输出: "eert"
解释: 'e'出现两次，'r'和't'都只出现一次。
因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
示例 2:

输入: s = "cccaaa"
输出: "cccaaa"
解释: 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
注意"cacaca"是不正确的，因为相同的字母必须放在一起。
示例 3:

输入: s = "Aabb"
输出: "bbAa"
解释: 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
注意'A'和'a'被认为是两种不同的字符。

https://leetcode.cn/problems/sort-characters-by-frequency/
*/
func frequencySort(s string) string {
	m := make(map[byte]int, len(s))
	for i := range s {
		m[s[i]]++
	}
	type pair struct {
		char byte
		cnt  int
	}
	pairs := make([]pair, 0, len(m))
	for k, v := range m {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].cnt > pairs[j].cnt
	})

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.char}, p.cnt)...)
	}

	return string(ans)
}

// 桶排序
func frequencySort2(s string) string {
	m := make(map[byte]int, len(s))
	maxCount := 0
	for i := range s {
		m[s[i]]++
		maxCount = max(maxCount, m[s[i]])
	}
	bucket := make([][]byte, maxCount+1)
	for k, v := range m {
		bucket[v] = append(bucket[v], k)
	}
	ans := make([]byte, 0, len(s))
	for i := maxCount; i > 0; i-- {
		for _, ch := range bucket[i] {
			ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
		}
	}

	return string(ans)
}
