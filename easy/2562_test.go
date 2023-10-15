package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findTheArrayConcVal(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"",
			args{[]int{7, 52, 2, 4}},
			596,
		},
		{
			"",
			args{[]int{5, 14, 13, 8, 12}},
			673,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTheArrayConcVal(tt.args.nums), "findTheArrayConcVal(%v)", tt.args.nums)
		})
	}
}

func Test_getUnion(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{1, 2},
			12,
		},
		{
			"",
			args{1, 20},
			120,
		},
		{
			"",
			args{11, 22},
			1122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getUnion(tt.args.num1, tt.args.num2), "getUnion(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
