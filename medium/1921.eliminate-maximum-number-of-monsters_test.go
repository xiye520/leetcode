package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_eliminateMaximum(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 3, 4}, []int{1, 1, 1}},
			3,
		},
		{
			"",
			args{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}},
			1,
		},
		{
			"",
			args{[]int{3, 2, 4}, []int{5, 2, 3}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, eliminateMaximum(tt.args.dist, tt.args.speed), "eliminateMaximum(%v, %v)", tt.args.dist, tt.args.speed)
		})
	}
}

func Test_eliminateMaximum2(t *testing.T) {
	type args struct {
		dist  []int
		speed []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{[]int{1, 3, 4}, []int{1, 1, 1}},
			3,
		},
		{
			"",
			args{[]int{1, 1, 2, 3}, []int{1, 1, 1, 1}},
			1,
		},
		{
			"",
			args{[]int{3, 2, 4}, []int{5, 2, 3}},
			1,
		},
		{
			"",
			args{[]int{4, 2, 3}, []int{2, 1, 1}},
			3,
		},
		{
			"4,3,3,3,4",
			args{[]int{4, 3, 3, 3, 4}, []int{1, 1, 1, 1, 4}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, eliminateMaximum2(tt.args.dist, tt.args.speed), "eliminateMaximum2(%v, %v)", tt.args.dist, tt.args.speed)
		})
	}
}
