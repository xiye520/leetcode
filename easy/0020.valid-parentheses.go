package easy

/*
20. 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
func isValid(s string) bool {
	size := len(s)
	if size%2 != 0 {
		return false
	}
	top := 0
	stack := make([]int32, size/2+1) // 用切片做栈存储字符，每次将切片的最后一个弹出来
	for i, c := range s {

		switch s[i] {
		case '(':
			stack[top] = c + 1 // ')' 转int后等于'(' + 1
			top++
		case '[', '{':
			stack[top] = c + 2 // ']', '}'转int后等于 '[', '{'  + 2
			top++
		case ')', ']', '}':
			if top > 0 && stack[top-1] == c {
				top--
			} else {
				return false
			}
		}
		if top > size/2 { //符号开闭不等，必为失败
			return false
		}
	}

	return top == 0 // 最终的栈要为零
}

/*
结题思路：
1.用一个切片c保存未匹配的字符
2.symbol用来保存右括号与左括号的对应关系，因为只有右括号出现的时候才可能匹配到左括号，从切片c中出栈
3.判断条件：栈内有字符&&当前字符与栈顶字符匹配
*/
func isValid2(s string) bool {
	var c []byte
	symbol := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, value := range s {
		clen := len(c)
		if clen > 0 {
			if _, ok := symbol[byte(value)]; ok {
				if c[clen-1] == symbol[byte(value)] {
					c = c[:clen-1]
					continue
				}
			}
		}
		c = append(c, byte(value))
	}
	return len(c) == 0
}

func isValid3(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, len(s)/2+1)
	top := 0
	for _, char := range s {
		switch char {
		case '[', '{':
			stack[top] = char + 2
			top++
		case '(':
			stack[top] = char + 1
			top++
		case ')', ']', '}':
			if top <= 0 || stack[top-1] != char {
				return false
			}
			top--
		default:
			return false
		}
		if top > len(s)/2 {
			return false
		}
	}

	return top == 0
}
