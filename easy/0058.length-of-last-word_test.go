package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Hello World",
			args{"Hello World"},
			5,
		},
		{
			"   fly me   to   the moon  ",
			args{"   fly me   to   the moon  "},
			4,
		},
		{
			"luffy is still joyboy",
			args{"luffy is still joyboy"},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLastWord(tt.args.s), "lengthOfLastWord(%v)", tt.args.s)
		})
	}
}
