package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid3(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"()",
			args{"()"},
			true,
		},
		{
			"()[]{}",
			args{"()[]{}"},
			true,
		},
		{
			"(]",
			args{"(]"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isValid3(tt.args.s), "isValid3(%v)", tt.args.s)
		})
	}
}
