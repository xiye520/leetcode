package util

import "testing"

/*
Binary		二进制
Octal		八进制
decimal		十进制
Hexadecimal	十六进制
*/

func Test_Bin(t *testing.T) {
	a := BinDec("11111111111111111111111111111111")
	t.Log(a)
}
