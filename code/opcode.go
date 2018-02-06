package code

type OpCode int32

const (
	NOP = iota

	ADD
	DEC
	EX
	INC
	LD
	RLCA
	RRCA
)
