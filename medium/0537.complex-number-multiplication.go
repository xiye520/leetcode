package medium

import (
	"fmt"
	"strconv"
	"strings"
)

/**
复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：

实部 是一个整数，取值范围是 [-100, 100]
虚部 也是一个整数，取值范围是 [-100, 100]
i2 == -1
给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。

示例 1：

输入：num1 = "1+1i", num2 = "1+1i"
输出："0+2i"
解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。
示例 2：

输入：num1 = "1+-1i", num2 = "1+-1i"
输出："0+-2i"
解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
提示：

num1 和 num2 都是有效的复数表示。
*/
func complexNumberMultiply(num1 string, num2 string) string {
	real1, imag1 := parseNum(num1)
	real2, imag2 := parseNum(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+real2*imag1)
}

// 解析输入的数字，分别取出实部和虚部
func parseNum(num string) (int, int) {

	index := strings.Index(num, "+")
	// +号之前的全是实部
	real, _ := strconv.Atoi(num[:index])
	imag, _ := strconv.Atoi(num[index+1 : len(num)-1])

	return real, imag
}
