package code

import (
	"encoding/binary"
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
	entry := decodeTable[b]
	if entry.cont != nil {
		entry.cont(r, &entry.proto)
	}
	return entry.proto, nil
}

type decodeContinuation = func(io.ByteReader, *Instruction) error

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

func decodeInt16Source(r io.ByteReader, inst *Instruction) (err error) {
	*inst.Src(), err = decodeInt16(r)
	return
}

func decodeInt16(r io.ByteReader) (Operand, error) {
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
