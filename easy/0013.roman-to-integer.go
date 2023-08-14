package easy

/**
13. 罗马数字转整数
https://leetcode.cn/problems/roman-to-integer/

示例1:

输入:s = "III"
输出: 3
示例2:

输入:s = "IV"
输出: 4
示例3:

输入:s = "IX"
输出: 9
示例4:

输入:s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例5:

输入:s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

*/
var roman = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum = 0
	var now, next int
	for i := 0; i < len(s); i++ {
		now = roman[s[i]]
		if i == len(s)-1 {
			sum += now
			break
		}
		next = roman[s[i+1]]
		if now < next {
			// 下一位数字比当前的大，代表需要减位
			sum += (next - now)
			i++
		} else {
			sum += now
		}
	}

	return sum
}
