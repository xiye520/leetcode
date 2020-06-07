package easy

/*
1221. 分割平衡字符串

在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。
给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
返回可以通过分割得到的平衡字符串的最大数量。

示例 1：
输入：s = "RLRRLLRLRL"
输出：4
解释：s 可以分割为 "RL", "RRLL", "RL", "RL", 每个子字符串中都包含相同数量的 'L' 和 'R'。

示例 2：
输入：s = "RLLLLRRRLR"
输出：3
解释：s 可以分割为 "RL", "LLLRRR", "LR", 每个子字符串中都包含相同数量的 'L' 和 'R'。

示例 3：
输入：s = "LLLLRRRR"
输出：1
解释：s 只能保持原样 "LLLLRRRR".

提示：
1 <= s.length <= 1000
s[i] = 'L' 或 'R'
*/
func balancedStringSplit(s string) int {
	// 由于原串已经是平衡的了，我们遍历原串，记录当前 LL 和 RR 的数量差，如果数量差为 00 就表示可以以当前位置作为一个分割点。
	count := 0
	l, r := 0, 0
	for _, v := range s {
		if v == 'L' {
			l++
		} else {
			r++
		}

		if l > 0 && l == r {
			count++
			l = 0
			r = 0
			continue
		}
	}

	return count
}
