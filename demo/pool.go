package main

import "log"

/**
一个有意思的题目，求最大蓄水面积
简单介绍一下，找到两个柱子，使其“蓄水”面积最大譬如

Input: [1,8,6,2,5,4,8,3,7]
Output: 49

*/
func maxArea(height []int) (result int) {
	for i, j := 0, len(height)-1; i < j; {
		now := (j - i) * minInt(height[i], height[j])
		result = maxInt(result, now)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return result
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		log.Fatalf("maxArea not match 49,it is: %d\n", result)
	}
	log.Printf("success maxArea: %d", result)
}
