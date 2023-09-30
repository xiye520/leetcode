package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"",
			args{"/home/"},
			"/home",
		},
		{
			"",
			args{"/../"},
			"/",
		},
		{
			"",
			args{"/home//foo/"},
			"/home/foo",
		},
		{
			"",
			args{"/a/./b/../../c/"},
			"/c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, simplifyPath(tt.args.path), "simplifyPath(%v)", tt.args.path)
		})
	}
}
