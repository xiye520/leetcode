package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
			6,
		},
		{
			"",
			args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestOnes(tt.args.nums, tt.args.k), "longestOnes(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}

func Test_longestOnes2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2},
			6,
		},
		{
			"",
			args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestOnes2(tt.args.nums, tt.args.k), "longestOnes(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
