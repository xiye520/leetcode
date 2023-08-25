package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canCompleteCircuit2(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			3,
		},
		{
			"",
			args{[]int{2, 3, 4}, []int{3, 4, 3}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canCompleteCircuit2(tt.args.gas, tt.args.cost), "canCompleteCircuit2(%v, %v)", tt.args.gas, tt.args.cost)
		})
	}
}
