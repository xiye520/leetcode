package medium

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		for j := len(s); j >= i; j-- {
			if isBalanceStr(s[i:j]) && len(s[i:j]) > len(res) {
				res = s[i:j]
			}
		}
	}

	return res
}

func isBalanceStr(s string) bool {
	if len(s) <= 1 {
		return true
	}

	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
