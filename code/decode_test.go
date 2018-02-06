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
		{
			name:     "0x02 => ld (bc), a",
			bytes:    []byte{0x02},
			expected: NewBinaryInstruction(LD, BC_INDIR, A),
		},
		{
			name:     "0x03 => inc bc",
			bytes:    []byte{0x03},
			expected: NewUnaryInstruction(INC, BC),
		},
		{
			name:     "0x04 => inc b",
			bytes:    []byte{0x04},
			expected: NewUnaryInstruction(INC, B),
		},
		{
			name:     "0x05 => dec b",
			bytes:    []byte{0x05},
			expected: NewUnaryInstruction(DEC, B),
		},
		{
			name:     "0x06 => ld b, *",
			bytes:    []byte{0x06, 0x42},
			expected: NewBinaryInstruction(LD, B, 0x42),
		},
		{
			name:     "0x07 => rlca",
			bytes:    []byte{0x07},
			expected: NewNularyInstruction(RLCA),
		},
		{
			name:     "0x08 => ex af, af'",
			bytes:    []byte{0x08},
			expected: NewBinaryInstruction(EX, AF, AF_PRIME),
		},
		{
			name:     "0x09 => add hl, bc",
			bytes:    []byte{0x09},
			expected: NewBinaryInstruction(ADD, HL, BC),
		},
		{
			name:     "0x0a => ld a, (bc)",
			bytes:    []byte{0x0a},
			expected: NewBinaryInstruction(LD, A, BC_INDIR),
		},
		{
			name:     "0x0b => dec bc",
			bytes:    []byte{0x0b},
			expected: NewUnaryInstruction(DEC, BC),
		},
		{
			name:     "0x0c => inc c",
			bytes:    []byte{0x0c},
			expected: NewUnaryInstruction(INC, C),
		},
		{
			name:     "0x0d => dec c",
			bytes:    []byte{0x0d},
			expected: NewUnaryInstruction(DEC, C),
		},
		{
			name:     "0x0e => ld c, *",
			bytes:    []byte{0x0e, 0x6f},
			expected: NewBinaryInstruction(LD, C, 0x6f),
		},
		{
			name:     "0x0f => rrca",
			bytes:    []byte{0x0f},
			expected: NewNularyInstruction(RRCA),
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
