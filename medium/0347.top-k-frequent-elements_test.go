package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1,1,1,2,2,3",
			args{[]int{1, 1, 1, 2, 2, 3}, 2},
			[]int{1, 2},
		},
		{
			"1",
			args{[]int{1}, 1},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topKFrequent(tt.args.nums, tt.args.k), "topKFrequent(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
