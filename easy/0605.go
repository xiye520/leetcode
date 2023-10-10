package easy

/*
* 605. 种花问题
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，
能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。

示例 1：
输入：flowerbed = [1,0,0,0,1], n = 1
输出：true

示例 2：
输入：flowerbed = [1,0,0,0,1], n = 2
输出：false

https://leetcode.cn/problems/can-place-flowers/
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	l := len(flowerbed)
	sum := 0
	if l == 1 {
		sum = 1 - flowerbed[0]
		return sum >= n
	}
	if flowerbed[1] == 0 && flowerbed[0] == 0 {
		sum++
		flowerbed[0] = 1
	}

	for i := 1; i < l-1; i++ {
		if flowerbed[i] == 0 {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				sum++
				flowerbed[i] = 1
			}
			if sum >= n {
				return true
			}
		}
	}
	if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
		sum++
	}

	return sum >= n
}

func canPlaceFlowers3(flowerbed []int, n int) bool {
	flowerbed = append(append([]int{0}, flowerbed...), 0)
	for i := 1; i < len(flowerbed)-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1 // 种花！
			n--
		}
	}
	return n <= 0
}

func canPlaceFlowers4(flowerbed []int, n int) bool {
	for i, v := range flowerbed {
		if v == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}
