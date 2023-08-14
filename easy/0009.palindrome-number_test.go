package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 9.回文数字
func Test_isPalindrome(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: 12321,
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: 1231,
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: -12321,
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isPalindrome2(p.one), "输入:%v", p)
	}
}
