package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{"11", "1"},
			"100",
		},
		{
			"10101",
			args{"1010", "1011"},
			"10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addBinary(tt.args.a, tt.args.b), "addBinary(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
