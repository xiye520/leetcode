package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sumOfMultiples(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{7},
			21,
		},
		{
			"",
			args{10},
			40,
		},
		{
			"",
			args{9},
			30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sumOfMultiples(tt.args.n), "sumOfMultiples(%v)", tt.args.n)
		})
	}
}
