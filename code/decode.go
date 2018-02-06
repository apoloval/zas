package code

import (
	"encoding/binary"
	"io"
)

// Decoder decodes bytes into code instructions
type Decoder interface {
	Decode(r io.ByteReader, inst *Instruction) (nread int, err error)
}

// DecoderFunc implements Decoder for functions
type DecoderFunc func(r io.ByteReader, inst *Instruction) (nread int, err error)

// MachineDecoder decodes bytes into code instructions following the Z80 machine encoding.
type MachineDecoder struct{}

// Decode bytes into instructions following the Z80 machine encoding.
func (d MachineDecoder) Decode(r io.ByteReader, inst *Instruction) (nread int, err error) {
	b, err := r.ReadByte()
	if err != nil {
		return 1, err
	}
	entry := decodeTable[b]
	*inst = entry.proto
	if entry.cont != nil {
		n, err := entry.cont(r, inst)
		return n + 1, err
	}
	return 1, nil
}

type decodeContinuation = func(io.ByteReader, *Instruction) (nread int, err error)

var decodeTable = []struct {
	proto Instruction
	cont  decodeContinuation
}{
	// 0x00: nop
	{
		proto: NewNularyInstruction(NOP),
	},
	// 0x01: ld bc, **
	{
		proto: NewBinaryInstruction(LD, BC, 0),
		cont:  decodeInt16Source,
	},
}

func decodeInt16Source(r io.ByteReader, inst *Instruction) (int, error) {
	return decodeInt16(r, inst.Src())
}

func decodeInt16(r io.ByteReader, op *Operand) (int, error) {
	b0, err := r.ReadByte()
	if err != nil {
		return 1, err
	}
	b1, err := r.ReadByte()
	if err != nil {
		return 2, err
	}
	buf := []byte{b0, b1}
	*op = Operand(binary.LittleEndian.Uint16(buf))
	return 2, nil
}
