package main

import "fmt"

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
