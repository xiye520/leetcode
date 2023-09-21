package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_change(t *testing.T) {
	type args struct {
		amount int
		coins  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{5, []int{1, 2, 5}},
			4,
		},
		{
			"",
			args{3, []int{2}},
			0,
		},
		{
			"",
			args{10, []int{10}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, change(tt.args.amount, tt.args.coins), "change(%v, %v)", tt.args.amount, tt.args.coins)
		})
	}
}
