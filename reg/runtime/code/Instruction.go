package code

import (
	"fmt"
)

// An Instruction compiles all the information an instruction can hold
// at execution time
type Instruction struct {
	Line int
	Code Bytecode
	RA   byte
	RB   byte
	RC   byte
}

// Make constructs a new Instruction object from a raw uint32 with this template
//
// 0x FF FF FF FF
// -- RC RB RA BC
//
// Where RA, RB, RC are the three operands and BC is the Bytecode instruction
func Make(line int, raw uint32) Instruction {
	rc := (raw & 0xFF000000) >> 24
	rb := (raw & 0x00FF0000) >> 16
	ra := (raw & 0x0000FF00) >> 8
	bc := (raw & 0x000000FF) >> 0
	return Instruction{line, Bytecode(bc), byte(ra), byte(rb), byte(rc)}
}

func (i *Instruction) String() string {
	return fmt.Sprintf("%14s   %1d %1d %1d", i.Code, i.RA, i.RB, i.RC)
}
