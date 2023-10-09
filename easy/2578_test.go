package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_splitNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{4325},
			59,
		},
		{
			"",
			args{687},
			75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, splitNum(tt.args.num), "splitNum(%v)", tt.args.num)
		})
	}
}

func Test_splitNum2(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{4325},
			59,
		},
		{
			"",
			args{687},
			75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, splitNum2(tt.args.num), "splitNum2(%v)", tt.args.num)
		})
	}
}
