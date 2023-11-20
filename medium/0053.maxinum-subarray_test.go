package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
		{
			"",
			args{[]int{1}},
			1,
		},
		{
			"",
			args{[]int{5, 4, -1, 7, 8}},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubArray(tt.args.nums), "maxSubArray(%v)", tt.args.nums)
		})
	}
}

func Test_maxSubArray2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			6,
		},
		{
			"",
			args{[]int{1}},
			1,
		},
		{
			"",
			args{[]int{5, 4, -1, 7, 8}},
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSubArray2(tt.args.nums), "maxSubArray2(%v)", tt.args.nums)
		})
	}
}
