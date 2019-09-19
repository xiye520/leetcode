package main

import "fmt"

// 数组中重复的数字
func topic3() {
	// 交换函数
	swap := func(list []int, i, j int) {
		list[i], list[j] = list[j], list[i]
	}

	input := []int{2, 3, 1, 0, 2, 5}
	for i := 0; i < len(input); i++ {
		// 将转换后
		for i != input[i] {

			// i要换到input[i]的位置，但是如果input[i]位置的值（intput[input[i]]）等与input[i]，说明重复来了
			if input[i] == input[input[i]] {
				fmt.Println(input[i], "重复")
				goto END
			}
			swap(input, i, input[i])
		}
	}
END:
	fmt.Println("done")
}

// 4.二维数组中的查找

// 我自己最初想到的方法
// 时间消费在算等差值d这个数组(这里懒得写)，以及下面的关于row的循环
func topic4() {
	list := [5][5]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	// 等差值,都为3只是特殊
	d := [5]int{3, 3, 3, 3, 3}

	// 行列数
	col := 5
	row := 5
	var input int
	fmt.Scanf("%d", &input)

	for i := 0; i < row; i++ {
		if input > list[i][col-1] || input < list[i][0] {
			// 比该行最大还大或者比最小的还小，下一行
			continue
		} else {
			if (input-list[i][0])%d[i] == 0 {
				fmt.Println("true")
				goto DONE
			}
			continue
		}
	}
	fmt.Println("false")
DONE:
	fmt.Println("done")
}
