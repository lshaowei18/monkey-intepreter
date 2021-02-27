package vm

import (
	"fmt"
	"monkey/m/v2/code"
	"monkey/m/v2/compiler"
	"monkey/m/v2/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int // StackPointer. Always points to the next value. Top of the stack is stack[sp-1]
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.Opcode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			err := vm.handleConstant(ip)
			if err != nil {
				return err
			}
			ip += 2
		case code.OpAdd:
			err := vm.handleAdd()
			if err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		}
	}
	return nil
}

func (vm *VM) handleConstant(ip int) error {
	constIndex := code.ReadUint16(vm.instructions[ip+1:])

	err := vm.push(vm.constants[constIndex])
	return err
}

func (vm *VM) handleAdd() error {
	right := vm.pop()
	left := vm.pop()

	leftValue := left.(*object.Integer).Value
	rightValue := right.(*object.Integer).Value

	result := leftValue + rightValue
	err := vm.push(&object.Integer{Value: result})
	return err
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) LastPoppedStackElem() object.Object {
	return vm.stack[vm.sp]
}
