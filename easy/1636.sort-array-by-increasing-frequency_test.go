package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"3,1,1,2,2,2",
			args{[]int{1, 1, 2, 2, 2, 3}},
			[]int{3, 1, 1, 2, 2, 2},
		},
		{
			"1,3,3,2,2",
			args{[]int{2, 3, 1, 3, 2}},
			[]int{1, 3, 3, 2, 2},
		},
		{
			"5,-1,4,4,-6,-6,1,1,1",
			args{[]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}},
			[]int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, frequencySort(tt.args.nums), "frequencySort(%v)", tt.args.nums)
		})
	}
}
