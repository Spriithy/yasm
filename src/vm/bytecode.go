package small

const (
	//
	// general purpose opcodes
	//

	// SWI (ra | imm)       interrupts the execution to perform the expected software interrupt
	SWI = 0x00

	// MOV ra, (rb | imm)   moves the content from rb or the immediate value into register ra
	MOV = 0x01

	// MOM ra, rb           moves in memory content from the address stored in rb to the address stored in ra
	MOM = 0x02

	// LOA ra, rb           load the value stored at the memory address of rb into ra
	LOA = 0x03

	// STR ra, (rb | imm)   stores the content of register rb or immediate value into ra
	STR = 0x04

	//
	// arithmetic operations, each one updates the rx register to the operation result
	//

	// ADD ra, (rb | imm)   adds the second operand to the first register
	ADD = 0x05

	// SUB ra, (rb | imm)   subtract the second operand from the first register
	SUB = 0x06

	// MUL ra, (rb | imm)   multiplies the second operand with the first register
	MUL = 0x07

	// DIV ra, (rb | imm)   divide the first register by the second operand
	DIV = 0x08

	// REM ra, (rb | imm)   adds the second operand to the first register
	REM = 0x09

	// BSL ra, (rb | imm)   bit shift left the first register by the second operand
	BSL = 0x0a

	// BSR ra, (rb | imm)   bit shift right the first register by the second operand
	BSR = 0x0b

	// INC ra, (rb | imm)   increments the value stored in ra by the second operand signed value
	INC = 0x0c

	// DEC ra, (rb | imm)   decrements the value stored in ra by the second operand signed value
	DEC = 0x0d

	// AND ra, (rb | imm)   bitwise and on first register and second operand
	AND = 0x0e

	// IOR ra, (rb | imm)   bitwise ior on first register and second operand
	IOR = 0x0f

	// XOR ra, (rb | imm)   bitwise xor on first register and second operand
	XOR = 0x10

	// NOT ra               bitwise not on register ra
	NOT = 0x11

	//
	// all conditional jump statements update the value of rx to the comparison result (except for jmp)
	// rx = 1 if condition is met, rx = 0 otherwise
	//

	// JMP imm              performs a relative jump, ie. PC += imm
	JMP = 0x12

	// JZ  imm              relative jump if rx == 0
	JZ = 0x13

	// JNZ imm              relative jump if rx != 0
	JNZ = 0x14

	// JEQ (ra | imm), imm  relative jump if ra == rx (or imm, likewise)
	JEQ = 0x15

	// JNE (ra | imm), imm  relative jump if ra != rx
	JNE = 0x16

	// JLT (ra | imm), imm  relative jump if ra <  rx
	JLT = 0x17

	// JLE (ra | imm), imm  relative jump if ra <= rx
	JLE = 0x18

	// JGT (ra | imm), imm  relative jump if ra >  rx
	JGT = 0x19

	// JGE (ra | imm), imm  relative jump if ra >= rx
	JGE = 0x1a

	// SRL ra, rb           subroutine link
	SRL = 0x1b

	// RET ra, rb           return from subroutine
	RET = 0x1c
)
