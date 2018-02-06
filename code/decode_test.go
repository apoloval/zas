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
			name:     "0x01 => ld bc, **",
			bytes:    []byte{0x01, 0xcd, 0xab},
			expected: NewBinaryInstruction(LD, BC, 0xabcd),
		},
	} {
		decoder := MachineDecoder{}
		t.Run(test.name, func(t *testing.T) {
			var actual Instruction
			nread, err := decoder.Decode(bytes.NewReader(test.bytes), &actual)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, len(test.bytes), nread)
		})
	}
}
