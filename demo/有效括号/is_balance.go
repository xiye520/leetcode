package main

import "fmt"

func isBalance(str string) bool {
	// {}[]()
	//if len(str)%2 != 0 {
	//	return false
	//}
	m := make(map[int]int32, 0)
	var n int
	for _, c := range str {
		switch c {
		case '[':
			n++
			m[n] = c + 2
		case '{':
			n++
			m[n] = c + 2
		case '(':
			n++
			m[n] = c + 1
		default:
			//fmt.Println(int(c), n, m)
			if n <= 0 {
				return false
			}
			if v, ok := m[n]; ok {
				if v != c {
					return false
				}
			} else {
				return false
			}
			n--
		}
	}
	if n != 0 {
		fmt.Println(n)
		return false
	}

	return true
}
