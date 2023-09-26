package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_passThePillow(t *testing.T) {
	type args struct {
		n    int
		time int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{4, 5},
			2,
		},
		{
			"",
			args{3, 2},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, passThePillow(tt.args.n, tt.args.time), "passThePillow(%v, %v)", tt.args.n, tt.args.time)
		})
	}
}
