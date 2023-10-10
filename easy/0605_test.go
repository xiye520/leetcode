package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{[]int{1, 0, 0, 0, 1}, 1},
			true,
		},
		{
			"",
			args{[]int{1, 0, 0, 0, 1}, 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers(tt.args.flowerbed, tt.args.n), "canPlaceFlowers(%v, %v)", tt.args.flowerbed, tt.args.n)
		})
	}
}

func Test_canPlaceFlowers2(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"",
			args{[]int{0}, 1},
			true,
		},
		{
			"",
			args{[]int{1}, 0},
			true,
		},
		{
			"",
			args{[]int{0, 0, 0}, 1},
			true,
		},
		{
			"",
			args{[]int{1, 0, 0, 0, 1}, 1},
			true,
		},
		{
			"",
			args{[]int{1, 0, 0, 0, 1}, 2},
			false,
		},
		{
			"",
			args{[]int{0, 0, 1, 0, 0, 0, 1, 0, 0}, 3},
			true,
		},
		{
			"",
			args{[]int{0, 0, 1, 0, 0, 0, 1, 0}, 2},
			true,
		},
		{
			"",
			args{[]int{0, 1, 0, 0, 0, 1, 0, 0}, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers2(tt.args.flowerbed, tt.args.n), "canPlaceFlowers2(%v, %v)", tt.args.flowerbed, tt.args.n)
		})
	}
}
