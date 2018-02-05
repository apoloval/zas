package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperandOutOfInt16Range(t *testing.T) {
	for _, v := range []Operand{
		BC,
	} {
		assert.Condition(t, func() bool { return v > 0xffff })
	}
}
