package code

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Decoder decodes bytes into code instructions
type Decoder interface {
	Decode(r io.ByteReader) (Instruction, error)
}

// DecoderFunc implements Decoder for functions
type DecoderFunc func(r io.ByteReader) (Instruction, error)

// MachineDecoder decodes bytes into code instructions following the Z80 machine encoding.
type MachineDecoder struct{}

// Decode bytes into instructions following the Z80 machine encoding.
func (d MachineDecoder) Decode(r io.ByteReader) (Instruction, error) {
	b, err := r.ReadByte()
	if err != nil {
		return Instruction{}, err
	}
	switch b {
	case 0x00:
		return NewNularyInstruction(NOP), nil
	case 0x01:
		imm, err := d.decodeInt16(r)
		if err != nil {
			return Instruction{}, err
		}
		return NewBinaryInstruction(LD, BC, imm), nil
	}
	return Instruction{}, fmt.Errorf("illegal instruction code in 0x%x", b)
}

func (MachineDecoder) decodeInt16(r io.ByteReader) (Operand, error) {
	buf := [2]byte{}
	var err error
	buf[0], err = r.ReadByte()
	if err != nil {
		return 0, err
	}
	buf[1], err = r.ReadByte()
	if err != nil {
		return 0, err
	}
	return Operand(binary.LittleEndian.Uint16(buf[:])), nil
}
