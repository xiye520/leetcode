package easy

import (
	"testing"
)

/*
	go test -test.bench=".*" 26_remove_test.go 26_remove-duplicates-from-sorted-array.go

C:\gopath\src\leetcode\easy>go test -test.bench=".*" 26_remove_test.go 26_remove-duplicates-from-sorted-array.go
goos: windows
goarch: 386
Benchmark_removeDuplicates-2            50000000                30.4 ns/op
Benchmark_removeDuplicates2-2           50000000                35.6 ns/op
PASS
ok      command-line-arguments  3.444s

*/

func Benchmark_removeDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < b.N; i++ {
		removeDuplicates(nums)
	}
}

func Benchmark_removeDuplicates2(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	for i := 0; i < b.N; i++ {
		removeDuplicates2(nums)
	}
}

func Test_removeDuplicates2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	index := removeDuplicates2(nums)
	t.Log(nums[:index])
}
