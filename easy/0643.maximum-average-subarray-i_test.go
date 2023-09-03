package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1, 12, -5, -6, 50, 3",
			args{[]int{1, 12, -5, -6, 50, 3}, 4},
			12.75,
		},
		{
			"",
			args{[]int{5}, 1},
			5.000000,
		},
		{
			"",
			args{[]int{-1}, 1},
			-1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxAverage(tt.args.nums, tt.args.k), "findMaxAverage(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
