package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"100,4,200,1,3,2",
			args{[]int{100, 4, 200, 1, 3, 2}},
			4,
		},
		{
			"0,3,7,2,5,8,4,6,0,1",
			args{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestConsecutive(tt.args.nums), "longestConsecutive(%v)", tt.args.nums)
		})
	}
}
