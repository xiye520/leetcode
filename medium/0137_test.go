package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
			args{[]int{2, 2, 3, 2}},
			3,
		},
		{
			"",
			args{[]int{0, 1, 0, 1, 0, 1, 99}},
			99,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, singleNumber(tt.args.nums), "singleNumber(%v)", tt.args.nums)
		})
	}
}
