package main

import (
	"testing"
)

// go test -v -test.run Test_fizzBuzz
func Test_fizzBuzz(t *testing.T) {
	r := fizzBuzz(15)
	t.Log(r)
	for _, v := range r {
		t.Log(v)
	}
}
