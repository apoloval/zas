package code

type Instruction struct {
	Operation OpCode
	Operands  [2]Operand
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
