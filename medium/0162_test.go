package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
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
			args{[]int{1, 2, 3, 1}},
			2,
		},
		{
			"",
			args{[]int{1, 2, 1, 3, 5, 6, 4}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findPeakElement(tt.args.nums), "findPeakElement(%v)", tt.args.nums)
		})
	}
}
