package easy

/*
461. 汉明距离

两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:
输入: x = 1, y = 4
输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置。
*/
func hammingDistance(x int, y int) int {
	res := x ^ y // 异或，两个值不相同，则异或结果为1
	return hanming(res)
}

func hanming(n int) int {
	var count int
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
