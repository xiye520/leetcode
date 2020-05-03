package easy

func addDigits(num int) int {
	for num > 9 {
		num = (num-1)%9 + 1
	}
	return num
}

func addDigits2(num int) int {
	if num%9 == 0 && num != 0 {
		return 9
	} else {
		num %= 9
	}
	return num
}

func addDigits3(num int) int {
	// https://leetcode-cn.com/problems/add-digits/solution/java-o1jie-fa-de-ge-ren-li-jie-by-liveforexperienc/
	return (num-1)%9 + 1
}
