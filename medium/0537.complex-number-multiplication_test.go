package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{"1+1i", "1+1i"},
			"0+2i",
		},
		{
			"test2",
			args{"1+-1i", "1+-1i"},
			"0+-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, complexNumberMultiply(tt.args.num1, tt.args.num2), "complexNumberMultiply(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
