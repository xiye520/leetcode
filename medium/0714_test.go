package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit714(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 3, 2, 8, 4, 9}, 2},
			8,
		},
		{
			"",
			args{[]int{1, 3, 7, 5, 10, 3}, 3},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxProfit714(tt.args.prices, tt.args.fee), "maxProfit714(%v, %v)", tt.args.prices, tt.args.fee)
		})
	}
}

func Test_maxProfit714_2(t *testing.T) {
	type args struct {
		prices []int
		fee    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 3, 2, 8, 4, 9}, 2},
			8,
		},
		{
			"",
			args{[]int{1, 3, 7, 5, 10, 3}, 3},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxProfit714_2(tt.args.prices, tt.args.fee), "maxProfit714_2(%v, %v)", tt.args.prices, tt.args.fee)
		})
	}
}
