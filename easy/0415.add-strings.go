package easy

import (
	"bytes"
	"fmt"
	"github.com/xiye520/leetcode/util"
	"strconv"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和
*/
func addStrings(num1 string, num2 string) string {
	var ret string
	length := util.Max(len(num1), len(num2))
	var carry bool
	for i := 0; i < length; i++ {
		sum := 0
		if i < len(num1) {
			//a, _ = strconv.Atoi(num1[len(num1)-i-1 : len(num1)-i])
			sum += int(num1[len(num1)-i-1]) - 48
		}

		if i < len(num2) {
			//b, _ = strconv.Atoi(num2[len(num2)-i-1 : len(num2)-i])
			sum += int(num2[len(num2)-i-1]) - 48
		}
		if carry {
			sum += 1
			carry = false
		}

		// 写入当前位
		ret = fmt.Sprintf("%d%s", sum%10, ret)
		if sum >= 10 {
			carry = true
		}
	}
	if carry {
		ret = fmt.Sprintf("1%s", ret)
	}

	//fmt.Println(addStrings("0", "0"))
	//fmt.Println(addStrings("0", "1"))
	//fmt.Println(addStrings("99", "1"))
	//fmt.Println(addStrings("144", "66"))
	return ret
}

func addStrings2(num1 string, num2 string) string {
	var buf bytes.Buffer
	l1 := len(num1)
	l2 := len(num2)

	carry := false
	var result []string

	for i, j := l1-1, l2-1; i >= 0 || j >= 0; {
		digit := 0
		if carry {
			digit++
			carry = false
		}

		if i >= 0 {
			digit += int(num1[i]) - 48
			i--
		}
		if j >= 0 {
			digit += int(num2[j]) - 48
			j--
		}

		if digit >= 10 {
			carry = true
		}
		result = append(result, strconv.Itoa(digit%10))
	}

	if carry {
		result = append(result, "1")
	}

	for i := range result {
		buf.WriteString(result[len(result)-1-i])
	}

	return buf.String()
}

func addStrings3(num1 string, num2 string) string {
	// 使用buffer写，比字符串拼接或者Sprintf效率高
	length := util.Max(len(num1), len(num2))
	var bf bytes.Buffer
	res := make([]string, 0, length+1)
	var carry bool
	for i := 0; i < length; i++ {
		sum := 0
		if i < len(num1) {
			sum += int(num1[len(num1)-i-1]) - 48
		}

		if i < len(num2) {
			sum += int(num2[len(num2)-i-1]) - 48
		}
		if carry {
			sum += 1
			carry = false
		}

		if sum >= 10 {
			carry = true
		}
		// 写入当前位
		res = append(res, strconv.Itoa(sum%10))
	}
	if carry {
		res = append(res, "1")
	}

	for i := range res {
		bf.WriteString(res[len(res)-1-i])
	}

	return bf.String()
}
