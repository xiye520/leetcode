package easy

import (
	"fmt"
	"math"
	"strconv"
)

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

1563847412
输出
2147483651
0
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
func reverse(x int) int {
	t := int64(x)
	isMinus := false
	if t < 0 {
		isMinus = true
		t = -1 * t
	}

	st := fmt.Sprintf("%d", t)
	tmp := []byte(st)
	for i := 0; i < len(st)/2; i++ {
		j := len(st) - 1 - i
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	r, _ := strconv.ParseInt(string(tmp), 10, 64)
	if isMinus {
		r = -1 * r
	}

	if r > 1<<31-1-1 || r < -1<<31 {
		return 0
	}

	return int(r)
}

func reverse2(x int) int {
	sign := 1

	// 处理负数
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		// 取出x的末尾
		temp := x % 10
		// 放入 res 的开头
		res = res*10 + temp
		// x 去除末尾
		x = x / 10
	}

	// 还原 x 的符号到 res
	res = sign * res

	// 处理 res 的溢出问题
	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}
