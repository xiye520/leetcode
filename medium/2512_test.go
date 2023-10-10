package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_topStudents(t *testing.T) {
	type args struct {
		positive_feedback []string
		negative_feedback []string
		report            []string
		student_id        []int
		k                 int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"",
			args{[]string{"smart", "brilliant", "studious"},
				[]string{"not"},
				[]string{"this student is studious", "the student is smart"},
				[]int{1, 2},
				2},
			[]int{1, 2},
		},
		{
			"",
			args{[]string{"smart", "brilliant", "studious"},
				[]string{"not"},
				[]string{"this student is not studious", "the student is smart"},
				[]int{1, 2},
				2},
			[]int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, topStudents(tt.args.positive_feedback, tt.args.negative_feedback, tt.args.report, tt.args.student_id, tt.args.k), "topStudents(%v, %v, %v, %v, %v)", tt.args.positive_feedback, tt.args.negative_feedback, tt.args.report, tt.args.student_id, tt.args.k)
		})
	}
}
