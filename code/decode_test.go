package code

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	for _, test := range []struct {
		name     string
		bytes    []byte
		expected Instruction
	}{
		{
			name:     "0x00 => nop",
			bytes:    []byte{0x00},
			expected: NewNularyInstruction(NOP),
		},
		{
			name:     "0x01 => ld bc, i16",
			bytes:    []byte{0x01, 0xcd, 0xab},
			expected: NewBinaryInstruction(LD, BC, 0xabcd),
		},
	} {
		decoder := MachineDecoder{}
		t.Run(test.name, func(t *testing.T) {
			actual, err := decoder.Decode(bytes.NewReader(test.bytes))
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
		})
	}
}
