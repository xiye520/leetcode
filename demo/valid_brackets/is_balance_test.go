package main

import "testing"

// go test -v -test.run Test_isBalance
func Test_isBalance(t *testing.T) {
	par := map[string]bool{
		"[]{}()": true,
		"[{}]()": true,
		"[{}])(": false,
		"[{}])":  false,
		"[":      false,
		"[]}]":   false,
		"":       true,
	}

	for k, v := range par {
		res := isBalance(k)
		if v != res {
			t.Fatalf("str: %s, want: %v,res: %v", k, v, res)
		}
	}

}
