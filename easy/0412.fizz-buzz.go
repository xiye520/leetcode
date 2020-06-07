package easy

import "strconv"

/*
412. Fizz Buzz

写一个程序，输出从 1 到 n 数字的字符串表示。
1. 如果 n 是3的倍数，输出“Fizz”；
2. 如果 n 是5的倍数，输出“Buzz”；
3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
示例：
n = 15,
返回:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/
func fizzBuzz(n int) []string {
	s := make([]string, n)
	var fb bool
	for i := 1; i <= n; i++ {
		fb = false
		if i%3 == 0 {
			s[i-1] = "Fizz"
			fb = true
		}
		if i%5 == 0 {
			s[i-1] = s[i-1] + "Buzz"
			fb = true
		}

		if !fb {
			s[i-1] = strconv.Itoa(i)
		}
	}

	return s
}

func fizzBuzz2(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FizzBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}
