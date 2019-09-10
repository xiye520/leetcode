package main

import "fmt"

//二分法查找
//切片s是升序的
//k为待查找的整数
//如果查到有就返回对应角标,
//没有就返回-1
func BinarySearch(s []int, k int) int {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		m := (lo + hi) >> 1
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m
		}
	}
	return -1
}

//测试：
func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(BinarySearch(s, 6))

}
