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

func decodeInt8Source(r io.ByteReader, inst *Instruction) (int, error) {
	return decodeInt8(r, inst.Src())
}

func decodeInt16Source(r io.ByteReader, inst *Instruction) (int, error) {
	return decodeInt16(r, inst.Src())
}

func decodeInt8(r io.ByteReader, op *Operand) (int, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 1, err
	}
	*op = Operand(b)
	return 1, nil
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
	// 0x02: ld (bc), a
	{
		proto: NewBinaryInstruction(LD, BC_INDIR, A),
	},
	// 0x03: inc bc
	{
		proto: NewUnaryInstruction(INC, BC),
	},
	// 0x04: inc b
	{
		proto: NewUnaryInstruction(INC, B),
	},
	// 0x05: dec b
	{
		proto: NewUnaryInstruction(DEC, B),
	},
	// 0x06: ld b, *
	{
		proto: NewBinaryInstruction(LD, B, 0),
		cont:  decodeInt8Source,
	},
	// 0x07: rlca
	{
		proto: NewNularyInstruction(RLCA),
	},
	// 0x08: ex af, af'
	{
		proto: NewBinaryInstruction(EX, AF, AF_PRIME),
	},
	// 0x09: add hl, bc
	{
		proto: NewBinaryInstruction(ADD, HL, BC),
	},
	// 0x0a: ld a, (bc)
	{
		proto: NewBinaryInstruction(LD, A, BC_INDIR),
	},
	// 0x0b: dec bc
	{
		proto: NewUnaryInstruction(DEC, BC),
	},
	// 0x0c: inc c
	{
		proto: NewUnaryInstruction(INC, C),
	},
	// 0x0d: dec c
	{
		proto: NewUnaryInstruction(DEC, C),
	},
	// 0x0e: ld c, *
	{
		proto: NewBinaryInstruction(LD, C, 0),
		cont:  decodeInt8Source,
	},
	// 0x0f: rrca
	{
		proto: NewNularyInstruction(RRCA),
	},
}
