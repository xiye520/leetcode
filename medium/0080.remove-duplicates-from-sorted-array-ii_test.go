package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1,1,1,2,2,3",
			args{[]int{1, 1, 1, 2, 2, 3}},
			5,
		},
		{
			"1,1,1,2,2,3",
			args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeDuplicates(tt.args.nums), "removeDuplicates(%v)", tt.args.nums)
		})
	}
}
