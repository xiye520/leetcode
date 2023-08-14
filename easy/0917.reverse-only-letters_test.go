package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseOnlyLetters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"success",
			args{s: "ab-cd"},
			"dc-ba",
		},
		{
			"success2",
			args{s: "a-bC-dEf-ghIj"},
			"j-Ih-gfE-dCba",
		},
		{
			"success3",
			args{s: "Test1ng-Leet=code-Q!"},
			"Qedo1ct-eeLg=ntse-T!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseOnlyLetters(tt.args.s), "reverseOnlyLetters(%v)", tt.args.s)
		})
	}
}
