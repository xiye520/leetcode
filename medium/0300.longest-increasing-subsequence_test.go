package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
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
			args{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			4,
		},
		{
			"",
			args{[]int{0, 1, 0, 3, 2, 3}},
			4,
		},
		{
			"",
			args{[]int{7, 7, 7, 7, 7, 7, 7}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLIS(tt.args.nums), "lengthOfLIS(%v)", tt.args.nums)
		})
	}
}
