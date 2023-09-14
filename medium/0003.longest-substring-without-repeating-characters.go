package medium

import (
	"math"
	"strings"
)

/*
3.无重复的字符的最长子串

给定一个字符串，找出不含有重复字符的最长子串的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 无重复字符的最长子串是 "abc"，其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 无重复字符的最长子串是 "b"，其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 无重复字符的最长子串是 "wke"，其长度为 3。
     请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
————————————————
版权声明：本文为CSDN博主「xvpenghao」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/heart66_A/article/details/83714168
*/

// abcabcbb
// 判断子串是否合法
// 内嵌一个循环判断是否有重复子串
// 如果有个重复子串，则start和end的位置增加
// 如果不重复则，继续增加子串的数量
func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	var strLen = len(s)
	var start, end = 0, 1
	var s2 string
	for end <= strLen {
		var s1 string = string(s[start:end])
		//计算重复
		var repeatCount int = 1
		var s1Len int = len(s1)
		for i := 0; i < s1Len; i++ {
			var temp string = string(s1[i])
			repeatCount = strings.Count(s1, temp)
			//说明有重复的了
			if repeatCount > 1 {
				start++
				end = start + len(s1)
				break
			}
		}
		if repeatCount == 1 {
			//增加
			s2 = s1 //保存之前的值
			end++
		}
	}
	return len(s2)
}

// 方法二 n*n*n  找出该串的所有子串，不重复，且最大即可
func lengthOfLongestSubstring2(s string) int {
	var counts int = 0
	for i := 0; i < len(s); i++ {
		//找出他所有的子串
		for j := i + 1; j <= len(s); j++ {
			if allUnique(s, i, j) {
				temp := math.Max(float64(counts), float64(j-i))
				counts = int(temp)
			}
		}
	}
	return counts
}

// 找出子串不重复的
func allUnique(s string, start, end int) bool {
	//会截取start -> end-1位置的下标
	s = s[start:end]
	for i := 0; i < len(s); i++ {
		//返回该子串在主串中的重复数量
		repeatCount := strings.Count(s, string(s[i]))
		if repeatCount != 1 {
			return false
		}
	}
	return true
}

// 方法三 利用滑动窗口和容器
func lengthOfLongestSubstring3(s string) int {
	//借助一个容器，来判断，子串中是否有重复 妙
	m := make(map[byte]byte)
	sLen := len(s)
	start, end := 0, 0
	var repeatCount int = 0
	//start 和 end 双条件判断，只有end一个也可以，可能这样更严谨一些吧
	for start < sLen && end < sLen {
		temp := s[end]
		if _, ok := m[temp]; !ok {
			//不存在说明该key是唯一的
			m[s[end]] = s[end]
			end++ //移动滑动窗口
			repeatCount = int(math.Max(float64(repeatCount), float64(end-start)))
		} else {
			//说明了有重复的，滑动窗口移动，则start+1,
			delete(m, s[start])
			start++
		}
	}
	return repeatCount

}

// 方法四 优化滑动窗口
// abcabcbb
// abc
// bca
// cab
// abc
// cb
// b
func lengthOfLongestSubstring4(s string) int {
	var n, ans = len(s), 0
	m := make(map[string]int) //存放字符出现的位置
	for j, i := 0, 0; j < n; j++ {
		if _, ok := m[string(s[j])]; ok {
			//发现重复的，则重新选择一个i，这个i停留在出现重复的前一位置
			i = max(m[string(s[j])], i) //这里的i则为下标
		}
		//每次计算j-i+1的记录，j为字符串的下标，i为下标，重复，i则从重复位置前一位开始
		ans = max(ans, j-i+1)
		m[string(s[j])] = j + 1
	}
	return ans
}

// 5.计算重复字符之间的最大长度，从而确定最长字符串
func lengthOfLongestSubstring5(s string) int {
	//输入只有ASCII字符，数组下标同ASCII码
	location := [256]int{}
	for i := range location {
		location[i] = -1
	}
	maxLen, left := 0, 0
	//最长无重复子字符串
	for i := 0; i < len(s); i++ {
		//确定未重复字符位置及最长位置
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}
	return maxLen
}

func lengthOfLongestSubstring6(s string) int {
	ans := 0
	m := make(map[uint8]bool, len(s))
	end := 0
	for i := 0; i < len(s); i++ {

		if i != 0 {
			// 重复
			delete(m, s[i-1])
		}
		for end < len(s) && !m[s[end]] {
			m[s[end]] = true
			end++
		}

		ans = max(ans, end-i)
	}

	return ans
}
