package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"7,1,5,3,6,4",
			args{[]int{7, 1, 5, 3, 6, 4}},
			7,
		},
		{
			"1,2,3,4,5",
			args{[]int{1, 2, 3, 4, 5}},
			4,
		},
		{
			"7,6,4,3,1",
			args{[]int{7, 6, 4, 3, 1}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxProfit(tt.args.prices), "maxProfit(%v)", tt.args.prices)
		})
	}
}
