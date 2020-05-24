package main

import (
	"testing"
)

// go test -v -test.run Test_dispatchCoin
func Test_dispatchCoin(t *testing.T) {
	left := dispatchCoin()
	t.Logf("剩下：%d", left)
}
