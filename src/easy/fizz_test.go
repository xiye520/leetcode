package easy

import (
	"strconv"
	"testing"
)

func FizzBuzz(n int) []string {
	var r []string
	for i := 1; i <= n; i++ {
		var a, b bool
		var s string
		if i%3 == 0 {
			s = "Fizz"
			a = true
		}
		if i%5 == 0 {
			if a {
				s = "FizzBuzz"
			} else {
				s = "Buzz"
			}
			b = true
		}
		if !a && !b {
			s = strconv.Itoa(i)
		}
		r = append(r, s)
	}

	return r
}

func TestFizzBuzz(t *testing.T) {
	r := FizzBuzz(15)
	for _, v := range r {
		t.Log(v)
	}
}
