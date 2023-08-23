package _821

import (
	"io/ioutil"
	"log"
)

/*
有二进制文件input.data,将其每一个字节都视为一个8位无符号整数,
编写代码，找出其中出现次数最多(可能并列)的值,输出其数值和出现次数,需注重代码效率.
*/

func readFile() {
	data, err := ioutil.ReadFile("input.data")
	if err != nil {
		log.Println(err)
		return
	}
	input := []uint8(data)
	maxCount(input)
}

func maxCount(input []uint8) ([]uint8, int) {
	count := 0
	l := min(len(input), 1>>8)
	m := make(map[uint8]int, l)

	for i := range input {
		m[input[i]]++
		count = max(count, m[input[i]])
	}
	slice := make([]uint8, 0, l)
	for k, v := range m {
		if v == count {
			slice = append(slice, k)
		}
	}

	return slice, count
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
