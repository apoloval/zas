package code

type Instruction struct {
	Operation OpCode
	Operands  [2]Operand
}

// Dest returns the destination operand
func (inst *Instruction) Dest() *Operand {
	return &inst.Operands[0]
}

// Src  returns the source operand
func (inst *Instruction) Src() *Operand {
	return &inst.Operands[1]
}

// NewNularyInstruction creates a new instruction with no operands
func NewNularyInstruction(opCode OpCode) Instruction {
	return Instruction{Operation: opCode}
}

// NewUnaryInstruction creates a new instruction with a single operand
func NewUnaryInstruction(opCode OpCode, operand Operand) Instruction {
	return Instruction{Operation: opCode, Operands: [2]Operand{operand}}
}

// NewBinaryInstruction creates a new instruction with a pair of operands
func NewBinaryInstruction(opCode OpCode, operand1 Operand, operand2 Operand) Instruction {
	return Instruction{Operation: opCode, Operands: [2]Operand{operand1, operand2}}
}
