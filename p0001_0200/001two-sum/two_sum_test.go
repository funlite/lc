package ptwosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pt struct {
	sv []int
	tv int
}

type rt struct {
	rv []int
}

type question struct {
	p pt
	r rt
}

func Test_OK(t *testing.T) {
	a := assert.New(t)

	qs := []question{
		question{
			p: pt{
				sv: []int{3, 2, 4},
				tv: 6,
			},
			r: rt{
				rv: []int{1, 2},
			},
		},
	}

	for _, q := range qs {
		p, r := q.p, q.r
		a.Equal(r.rv, twoSum(p.sv, p.tv), "input:%v", p)

	}
}
