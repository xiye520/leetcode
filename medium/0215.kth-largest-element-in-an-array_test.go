package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
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
			args{[]int{3, 2, 1, 5, 6, 4}, 2},
			5,
		},
		{
			"",
			args{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			4,
		},
		{
			"",
			args{[]int{3, 3, 3, 3, 4, 3, 3, 3, 3}, 5},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findKthLargest3(tt.args.nums, tt.args.k), "findKthLargest(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
