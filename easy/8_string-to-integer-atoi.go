package easy

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

/*
	8. 字符串转换整数 (atoi)
*/
func MyAtoi(str string) int {
	newStr := strings.Replace(str, " ", "", -1) //去掉空格
	if len(newStr) == 0 {
		return 0
	}

	isNature := string(newStr[0]) == "+" //去掉空格后第一个字符为正号
	isMinus := string(newStr[0]) == "-"  //去掉空格后第一个字符为负号
	if len(newStr) >= 2 && (isNature || isMinus) && !unicode.IsDigit(rune(newStr[1])) {
		return 0
	}

	var nstr string
	var hasNum bool
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			if hasNum {
				break
			}
			continue
		}

		if string(str[i]) == "+" {
			if hasNum {
				break
			}
			continue
		}
		if string(str[i]) == "-" {
			if hasNum {
				break
			}
			continue
		}
		if unicode.IsDigit(rune(str[i])) {
			hasNum = true
			nstr += string(str[i])
			continue
		}
		break
	}

	n, err := strconv.Atoi(nstr)
	if err != nil {
		log.Println(err)
		log.Println(fmt.Sprintf("%t", err))
		log.Println(strings.Contains(err.Error(), "value out of range"))
		if strings.Contains(err.Error(), "value out of range") {
			if isMinus {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		return 0
	}

	if isMinus {
		n = -n
	}

	if n > math.MaxInt32 {
		return math.MaxInt32
	}
	if n < math.MinInt32 {
		return math.MinInt32
	}

	return n
}
