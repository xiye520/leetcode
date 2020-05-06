package medium

/*
365. 水壶问题
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)

输入: x = 3, y = 5, z = 4
输出: True
示例 2:

输入: x = 2, y = 6, z = 5
输出: False
*/
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 || z == x+y { // 条件 1
		return true
	} else if z > x+y || y == 0 { // 条件 2 和 3
		return false
	}

	// 迭代公式 2、 4、 6
	return canMeasureWater(y, x%y, z%y)
}

// 如果x和y的最大公约数为1的话，那么经过多次迭代之后，可以凑出来[1,x+y]区间的任何正整数。
//如果不为1，提取x和y的最大公约数g之后，参照上述可以凑出来[1, (x+y)/g]区间的任何正整数，
//但是结果空间为[1, (x+y)/g]的正整数*最大公约数g。
func canMeasureWater2(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x == 0 || y == 0 {
		return z == x || z == y
	}

	g := GCD(x, y)
	return z%g == 0 && z >= 0 && z <= (x+y)
}

// 求两数的最大公约数
func GCD(x, y int) int {
	p1, p2 := x, y
	if p2 > p1 {
		p1, p2 = y, x
	}
	if p1%p2 == 0 {
		return p2
	}
	return GCD(p1%p2, p2)
}

func canMeasureWater3(x int, y int, z int) bool {
	if x > y {
		x, y = y, x
	}

	nums := make(map[int]int, 0)
	nums[x+y] = 1
	nums[0] = 1

	k := 0
	getMap(nums, x, y, k)

	if nums[z] == 1 {
		return true
	}

	return false
}

func getMap(nums map[int]int, x, y, k int) {
	if nums[x+k] == 0 {
		nums[x+k] = 1
		if x+k < y {
			getMap(nums, x, y, x+k)
		}
	}

	if x-k > 0 && nums[x-k] == 0 {
		nums[x-k] = 1
		getMap(nums, x, y, x-k)
	}

	if nums[y+k] == 0 {
		nums[y+k] = 1
	}

	if y-k > 0 && nums[y-k] == 0 {
		nums[y-k] = 1
		getMap(nums, x, y, y-k)
	}
}
