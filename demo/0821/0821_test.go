package _821

import (
	"reflect"
	"testing"
)

func Test_maxCount(t *testing.T) {
	type args struct {
		input []uint8
	}
	tests := []struct {
		name  string
		args  args
		want  []uint8
		want1 int
	}{
		{
			"success",
			args{[]uint8{1, 2, 3, 1, 2}},
			[]uint8{1, 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := maxCount(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxCount() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("maxCount() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
