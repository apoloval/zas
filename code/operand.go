package code

type Operand int32

const (
	_ = 0xffff + iota

	A
	AF
	AF_PRIME
	B
	BC
	BC_INDIR
	C
	D
	E
	H
	HL
	L
)
