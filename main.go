package main

import (
	"fmt"
	"regexp"
)

func main() {
	strs := []string{"   -42", "4193 with words", "words and 98", "-91283472332", "+", "-", "3.14159", "+1", "+-5", "+5-", "   +0 123",
		"20000000000000000000", "-   234", "+   234"}
	reg := regexp.MustCompile(".+?([0-9]+).+?")
	for _, str := range strs {
		result := reg.FindAllString(str, -1)
		fmt.Printf("'%s',result: %v\n", str, result)
		//if len(result) > 0 {
		//fmt.Printf("'%s',first: '%s'\n", str, result[0])
		//}
	}
}
