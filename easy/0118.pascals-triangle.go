package easy

import "fmt"

/*
118.杨辉三角

给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	a := make([][]int, 0, numRows)
	a = append(a, []int{1})
	for i := 2; i <= numRows; i++ {
		tmp := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				tmp[j] = 1
				continue
			}
			tmp[j] = a[i-2][j-1] + a[i-2][j]
		}
		a = append(a, tmp)
	}

	return a
}

func generate2(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	a := make([][]int, 0)
	a = append(a, []int{1})
	for i := 1; i < numRows; i++ {
		temp := []int{1}
		for j := 1; j <= i-1; j++ {
			temp = append(temp, a[i-1][j-1]+a[i-1][j])
		}
		temp = append(temp, 1)
		a = append(a, temp)
	}
	return a
}

func test_generate() {
	r := generate(10)
	//r := generate2(10)
	for _, v := range r {
		fmt.Println(v)
	}
}
