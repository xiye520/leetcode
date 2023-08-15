package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
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
			"ahbgdc",
			args{"abc", "ahbgdc"},
			true,
		},
		{
			"ahbgdc",
			args{"axc", "ahbgdc"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isSubsequence(tt.args.s, tt.args.t), "isSubsequence(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
