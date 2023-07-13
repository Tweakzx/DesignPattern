package interpreter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpret(t *testing.T) {
	//2 + 3 * 4 - 5 + 6/3 = 11
	exp := NewAddExpression(
		NewNumberExpression(2),
		NewAddExpression(
			NewSubExpression(
				NewMulExpression(
					NewNumberExpression(3),
					NewNumberExpression(4),
				),
				NewNumberExpression(5),
			),
			NewDivExpression(
				NewNumberExpression(6),
				NewNumberExpression(3),
			),
		),
	)

	assert.Equal(t, 11, exp.Interpret())
}
