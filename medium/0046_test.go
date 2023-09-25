package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"",
			args{[]int{1, 2, 3}},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			"",
			args{[]int{0, 1}},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			"",
			args{[]int{1}},
			[][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, permute(tt.args.nums), "permute(%v)", tt.args.nums)
		})
	}
}

func Test_permute2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"",
			args{[]int{1, 2, 3}},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
		{
			"",
			args{[]int{0, 1}},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			"",
			args{[]int{1}},
			[][]int{{1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, permute3(tt.args.nums), "permute2(%v)", tt.args.nums)
		})
	}
}
