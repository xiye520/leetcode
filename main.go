package leetcode

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("hello")

	str := "-056546161321615612315"
	n, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("str: %s转换为数字后为：%d\n", str, n)
}
