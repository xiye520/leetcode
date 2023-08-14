package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"egg, add",
			args{"egg", "add"},
			true,
		},
		{
			"foo, bar",
			args{"foo", "bar"},
			false,
		},
		{
			"paper, title",
			args{"paper", "title"},
			true,
		},
		{
			"badc, baba",
			args{"badc", "baba"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isIsomorphic(tt.args.s, tt.args.t), "isIsomorphic(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
