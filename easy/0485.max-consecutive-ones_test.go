package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
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
			args{[]int{1, 1, 0, 1, 1, 1}},
			3,
		},
		{
			"",
			args{[]int{1, 0, 1, 1, 0, 1}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaxConsecutiveOnes(tt.args.nums), "findMaxConsecutiveOnes(%v)", tt.args.nums)
		})
	}
}
