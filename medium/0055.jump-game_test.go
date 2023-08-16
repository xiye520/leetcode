package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canJump2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"2,3,1,1,4",
			args{[]int{2, 3, 1, 1, 4}},
			true,
		},
		{
			"3,2,1,0,4",
			args{[]int{3, 2, 1, 0, 4}},
			false,
		},
		{
			"0, 3, 1, 0, 4, 3, 2, 0, 1",
			args{[]int{0, 3, 1, 0, 4, 3, 2, 0, 1}},
			false,
		},
		{
			"0",
			args{[]int{0}},
			true,
		},
		{
			"0, 3, 1, 0, 4, 3, 2, 0, 1",
			args{[]int{2, 0, 0, 3, 0, 1, 4, 4}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canJump2(tt.args.nums), "canJump2(%v)", tt.args.nums)
		})
	}
}
