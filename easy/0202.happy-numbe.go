package easy

/*
输入：19
输出：true
解释：
1*1 + 9*9 = 82
8*8 + 2*2 = 68
6*6+ 8*8 = 100
1*1 + 0*0 + 0*0 = 1

*/

/*
方法：使用“快慢指针”思想找出循环：“快指针”每次走两步，“慢指针”每次走一步，当二者相等时，即为一个循环周期。
此时，判断是不是因为1引起的循环，是的话就是快乐数，否则不是快乐数。

注意：此题不建议用集合记录每次的计算结果来判断是否进入循环，因为这个集合可能大到无法存储；
另外，也不建议使用递归，同理，如果递归层次较深，会直接导致调用栈崩溃。不要因为这个题目给出的整数是int型而投机取巧。

链接：https://leetcode-cn.com/problems/happy-number/solution/shi-yong-kuai-man-zhi-zhen-si-xiang-zhao-chu-xun-h/
*/
func isHappy(n int) bool {
	slow, fast := n, n
	for true {
		slow = square(slow)
		fast = square(fast)
		fast = square(fast)
		if slow == fast {
			break
		}
	}

	return slow == 1
}

func square(n int) int {
	sum := 0
	for n > 0 {
		bit := n % 10
		sum += bit * bit
		n /= 10
	}

	return sum
}

// map解法
// https://leetcode-cn.com/problems/happy-number/solution/golangde-mapjie-fa-by-lryong/
func isHappy2(n int) bool {
	calHistory := make(map[int]struct{})
	if n <= 0 {
		return false
	}
	for n != 1 {
		if _, ok := calHistory[n]; ok {
			return false
		}
		calHistory[n] = struct{}{}
		n = square(n)
	}
	return true
}
