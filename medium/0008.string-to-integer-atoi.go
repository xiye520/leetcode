package medium

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

/*
 8. 字符串转换整数 (atoi)

请你来实现一个 atoi 函数，使其能将字符串转换成整数。

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，qing返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。

示例 1:

输入: "42"
输出: 42
示例 2:

输入: "   -42"
输出: -42
解释: 第一个非空白字符为 '-', 它是一个负号。

	我们尽可能将负号与后面所有连续出现的数字组合起来，最后得到 -42 。

示例 3:

输入: "4193 with words"
输出: 4193
解释: 转换截止于数字 '3' ，因为它的下一个字符不为数字。
示例 4:

输入: "words and 987"
输出: 0
解释: 第一个非空字符是 'w', 但它不是数字或正、负号。

	因此无法执行有效的转换。

示例 5:

输入: "-91283472332"
输出: -2147483648
解释: 数字 "-91283472332" 超过 32 位有符号整数范围。

	因此返回 INT_MIN (−231) 。
*/
func MyAtoi(str string) int {
	return myAtoi(str)
}

func myAtoi3(str string) int {

	return 0
}

/*
执行结果：pass
执行用时 :60 ms, 在所有 Go 提交中击败了5.65%的用户
内存消耗 :6.8 MB, 在所有 Go 提交中击败了5.35%的用户
*/
func myAtoi(str string) int {
	var firstIn, isMinus, getNum bool
	var ns string

	pattern := "\\d" //反斜杠要转义
	for _, b := range str {
		s := string(b)
		if !firstIn {
			if s != " " {
				firstIn = true
				//判断第一个非空字符类型
				switch s {
				case "-":
					isMinus = true
					continue
				case "+":
					continue
				default:
					result, _ := regexp.MatchString(pattern, s)
					if result {
						getNum = true
						ns += s
					} else {
						//第一个非空字符非法，返回0
						return 0
					}
				}
			}
			continue
		}

		//判断后续的是否是数字
		result, _ := regexp.MatchString(pattern, s)
		if result {
			getNum = true
			ns += s
		} else {
			break
		}
	}

	if !getNum {
		return 0
	}

	n, err := strconv.ParseInt(ns, 10, 64)
	if err != nil {
		if isMinus {
			return -1 << 31
		} else {
			return 1<<31 - 1
		}

	}

	if isMinus {
		n = -1 * n
	}

	if n > 1<<31-1 {
		return 1<<31 - 1
	} else if n < -1<<31 {
		return -1 << 31
	}

	return int(n)
}

func myAtoi2(s string) int {
	return convert(clean(s))
}

func clean(s string) (sign int, abs string) {
	// 除去首尾的空格
	s = strings.TrimSpace(s)
	if s == "" {
		return
	}

	switch s[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, s
	case '+':
		sign, abs = 1, s[1:]
	case '-':
		sign, abs = -1, s[1:]
	default:
		abs = ""
		return
	}

	for i, b := range abs {
		if b < '0' || '9' < b {
			abs = abs[:i]
			break
		}
	}

	return
}

func convert(sign int, absStr string) int {
	abs := 0

	for _, b := range absStr {
		abs = abs*10 + int(b-'0')
		// 检查溢出
		switch {
		case sign == 1 && abs > math.MaxInt32:
			return math.MaxInt32
		case sign == -1 && sign*abs < math.MinInt32:
			return math.MinInt32
		}
	}

	return sign * abs
}
