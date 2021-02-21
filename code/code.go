package code

import (
	"encoding/binary"
	"fmt"
)

type Instruction []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

func calculateInstructionLen(def *Definition) int {
	instructionLen := 1

	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	return instructionLen
}

func Make(op Opcode, operands ...int) []byte {
	def, ok := definitions[op]
	if !ok {
		return []byte{}
	}

	instructionLen := calculateInstructionLen(def)

	// Allocate byte slice with proper length
	instruction := make([]byte, instructionLen)
	instruction[0] = byte(op)

	offset := 1

	/*
		1. Iterate over the defined OeprandWidths
		2. Take matching element from operands & put it in the instruction
		3. Increase the offset
	*/
	for i, o := range operands {
		width := def.OperandWidths[i]

		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}
		offset += width
	}

	return instruction

}
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefinted", op)
	}

	return def, nil
}
