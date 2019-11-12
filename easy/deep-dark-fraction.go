package easy

/*
LDAP分式化简

示例 1：

输入：cont = [3, 2, 0, 2]
输出：[13, 4]
解释：原连分数等价于3 + (1 / (2 + (1 / (0 + 1 / 2))))。注意[26, 8], [-13, -4]都不是正确答案。
示例 2：

输入：cont = [0, 0, 3]
输出：[3, 1]
解释：如果答案是整数，令分母为1即可。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/deep-dark-fraction
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func fraction(cont []int) []int {
	res := make([]int, 2)
	res[0] = 1
	res[1] = 0
	for i := len(cont) - 1; i >= 0; i-- {
		tmp := res[1]
		res[1] = res[0]
		res[0] = cont[i]*res[1] + tmp
	}

	return res
}
