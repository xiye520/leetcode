package easy

/*
204. 计数质数
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
*/
func countPrimes(n int) int {
	//count := 0
	if n <= 2 {
		return 0
	}
	primaryArr := make([]int, 0, n)
	primaryArr = append(primaryArr, 2)
	for i := 3; i < n; i++ {
		if isPrime(i, primaryArr) {
			primaryArr = append(primaryArr, i)
		}
	}

	return len(primaryArr)
}

func isPrime(n int, primaryArr []int) bool {
	for _, v := range primaryArr {
		if n%v == 0 && n != v {
			//fmt.Println("---", n, v)
			return false
		}
		if v*v > n {
			break
		}
	}

	return true
}

func countPrimes1(n int) int {
	flag := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if !flag[i] {
			count++
			for j := 2; i*j < n; j++ {
				flag[i*j] = true
			}
		}
	}
	return count
}
