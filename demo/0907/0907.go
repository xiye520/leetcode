package _907

import (
	"fmt"
	"strings"
)

/**
一、 地址归一化问题
在葡语地址中，有一些单词表达的含义相同，如Rua 和 R、R. ，Avenida 和Av，Av.等，含义都表示街道，
现在给定一个同义词表，map<string, vector<string>>, key是街道的完整写法，value是对应的缩写列表，如Rua 对应R 和 R. 两个缩写。
现在给定两个地址 s1和s2，均由字母、数字、空格、逗号和. 组成，设计一个函数，判断s1中是否包含完整的s2的等价形式（不区分大小写）。
例如：

同义词表
"Avenida" ： "Av." , Av
"Rua" :          "R." , R
"Estrada":     "Estr."
"Alameda" :   "Al."
"Travessa" :   "Tv." , Tv
"Rodovia" :     "Rod."

s1 = Avenida Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.
s2 = Av. Jose Vasconcelos 345
return true

s1 = Av Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.
s2 = Avenida Jose Vasconcelos 345
return true

s1 = Av Jose Vasconcelos 345, Zona Santa Engracia, 66267 San Pedro Garza Garcia, N.L.
s2 = av. jose vasconcelos 345
return true

*/

func isSameAddress(m map[string][]string, s1, s2 string) bool {
	s1 = strings.Replace(s1, ",", "", -1)
	s2 = strings.Replace(s2, ",", "", -1)
	slice1 := strings.Fields(s1)
	slice2 := strings.Fields(s2)
	if len(slice1) < len(slice2) {
		return false
	}
	fmt.Println(slice1, len(slice1))
	fmt.Println(slice2, len(slice2))

	for i := range slice2 {
		if slice1[i] != slice2[i] {
			if !isMatchAddressMap(m, slice1[i], slice2[i]) && !isMatchAddressMap(m, slice2[i], slice1[i]) {
				return false
			}
		}
	}

	return true
}

func isMatchStr(m map[string][]string, s1, s2 string) bool {
	return isMatchAddressMap(m, s1, s2) || isMatchAddressMap(m, s2, s1)
}

func isMatchAddressMap(m map[string][]string, s1, s2 string) bool {
	if n := strings.Index(s2, "."); n > 0 {
		s1 = s1[:n]
		s2 = s2[:n]
	}
	if strings.ToLower(s1) == strings.ToLower(s2) {
		return true
	}

	addresses, ok := m[s1]
	if !ok {
		return false
	}

	for _, address := range addresses {
		if n := strings.Index(s2, "."); n > 0 {
			address = address[:n]
			s2 = s2[:n]
		}
		if strings.ToLower(address) == strings.ToLower(s2) {
			return true
		}
	}
	return false
}
