package main

import "strings"

// *匹配任意字符
// 当多条都匹配的时候，按前缀匹配最长的分发
// 前缀匹配+最长字串
func getCompareRoute(routes []string, url string) string {
	res := ""
	maxCount := 0
	for _, route := range routes {
		if strings.Contains(route, url) {
			return route
		}
		count := maxCompareCount(route, url)
		if count > maxCount {
			maxCount = count
			res = route
		}
	}

	return res
}

// *匹配任意字符，只有route才有*
func maxCompareCount(route, url string) int {
	count := 0
	for i := range route {
		if route[i] == '*' {
			// 进入最长字串匹配逻辑
			if i == len(route)-1 || count == len(url)-1 {
				return count + 1
			}
			// 找出route中*号之后，开头文本在url中的最长子串
			return count + 1 + getMaxPrefix(route[i+1:], url[count+1:])
		}
		if route[i] != url[count] {
			break
		}
		// 字符匹配，长度+1
		count++
	}

	return count
}

// 查看route在url中的最长公共前缀
func getMaxPrefix(route, url string) int {
	count := 0

	for i := range route {
		if url[i] == route[count] {
			count++
			if count >= len(url)-1 {
				break
			}
		}
	}
	return count
}
