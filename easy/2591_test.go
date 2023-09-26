package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_distMoney(t *testing.T) {
	type args struct {
		money    int
		children int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"",
			args{20, 3},
			1,
		},
		{
			"",
			args{16, 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, distMoney(tt.args.money, tt.args.children), "distMoney(%v, %v)", tt.args.money, tt.args.children)
		})
	}
}
