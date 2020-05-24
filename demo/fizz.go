package main

import (
	"strconv"
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

func fizzBuzz(n int) []string {
	s := make([]string, n)
	var fb bool
	for i := 1; i <= n; i++ {
		fb = false
		if i%3 == 0 {
			s[i-1] = "Fizz"
			fb = true
		}
		if i%5 == 0 {
			s[i-1] = s[i-1] + "Buzz"
			fb = true
		}

		if !fb {
			s[i-1] = strconv.Itoa(i)
		}
	}

	return s
}
