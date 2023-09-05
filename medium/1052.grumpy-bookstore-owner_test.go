package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3},
			16,
		},
		{
			"case 2",
			args{[]int{1}, []int{0}, 1},
			1,
		},
		{
			"case 3",
			args{[]int{3}, []int{1}, 1},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes), "maxSatisfied(%v, %v, %v)", tt.args.customers, tt.args.grumpy, tt.args.minutes)
		})
	}
}

func Test_maxSatisfied2(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3},
			16,
		},
		{
			"case 2",
			args{[]int{1}, []int{0}, 1},
			1,
		},
		{
			"case 3",
			args{[]int{3}, []int{1}, 1},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSatisfied2(tt.args.customers, tt.args.grumpy, tt.args.minutes), "maxSatisfied2(%v, %v, %v)", tt.args.customers, tt.args.grumpy, tt.args.minutes)
		})
	}
}
