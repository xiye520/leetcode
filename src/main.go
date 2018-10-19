package main

import "fmt"

func main() {
	Format("0x92c39294ca7f800001ca7f800001ca7f800001ca7f80000190")
	Format("0x92c39194ca3acccccdca00000000ca3f800000ca3f7fd27d")
}

func Format(str string) {
	var decodeStr string
	for i := 2; i < len(str); i++ {
		if i%2 == 0 {
			decodeStr += `\x`
		}

		decodeStr += string(str[i])
	}

	fmt.Println(decodeStr)
	fmt.Println("hello leetcode!")
}
