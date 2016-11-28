package small

import (
	"unsafe"
)

const (
	// Tail means the instruction is a 64-bits immediate value's tail
	Tail = 0x0

	// Head means the instruction is a 64-bits immediate value's head
	Head = 0x1
)

// Byte masks
const (
	xMask = 0x00000003
	tMask = 0x00000004
	oMask = 0x000000f8
	aMask = 0x00000f00
	bMask = 0x0000f000
	eMask = 0xffff0000
)

// An Instruction is a 32-bits unsigned integer that represents
// virtual-machine-executable informations.
type Instruction uint32

// ReadInstruction takes in 4 bytes and assembles them in a
// Instruction typed object.
func ReadInstruction(b0, b1, b2, b3 byte) Instruction {
	i := *new(Instruction)
	i |= Instruction(b0) << 0
	i |= Instruction(b1) << 8
	i |= Instruction(b2) << 16
	i |= Instruction(b3) << 24
	return i
}

// CompileInstruction is used to create an instruction from its description
func CompileInstruction(x, t, o, a, b uint32, e int16) Instruction {
	var i uint32
	i |= (x << 0) & xMask
	i |= (t << 2) & tMask
	i |= (o << 3) & oMask
	i |= (a << 8) & aMask
	i |= (b << 12) & bMask
	i |= (uint32(e) << 16) & eMask
	return Instruction(i)
}

// X returns the two meaningful bits that represent the Instruction
// length.
//      X = 0x00 -> usual one word instruction
//      X = 0x01 -> immediate one-word instruction (16-bits data)
//      X = 0x02 -> immediate two-word instruction (32-bits data)
// The last one is only valid on 64-bits machines:
//      X = 0x03 -> immediate three-word instruction (64-bits data)
func (i Instruction) X() uint32 {
	return uint32(i) & xMask
}

// T returns wether the instruction works on integers (or nothing)
// or floats.
// True     : Integers or nothing
// False    : Floats
func (i Instruction) T() bool {
	return (i&tMask)>>2 == 0x00
}

// O returns the byte value of the opcode that is stored in the instruction
func (i Instruction) O() byte {
	return byte((i & oMask) >> 3)
}

// RA returns the index of the targetted register
func (i Instruction) RA() int {
	return int(i & aMask >> 8)
}

// RB returns the index of the targetted register
func (i Instruction) RB() int {
	return int(i & bMask >> 12)
}

// E returns the extra 16-bits segment stored in the instruction
// This is mostly used as an immediate value
func (i Instruction) E() int16 {
	return int16(i >> 16)
}

// Int16 returns the int16 representation of the instruction's bits
// or rather, it's immediate value
func (i Instruction) Int16() int16 {
	return int16(i >> 16)
}

// Int32 returns the int32 representation of the instruction's bits
// or rather, it's immediate value
func (i Instruction) Int32() int32 {
	return int32(i >> 0)
}

// Float32 returns the float32 representation of the instruction's bits
// or rather, it's immediate value
func (i Instruction) Float32() float32 {
	return *(*float32)(unsafe.Pointer(&i))
}

// Int64 returns the int64 representation of the instruction's bits
// or rather, it's immediate value when passed another instruction
// interpreted as an int64
func (i Instruction) Int64(i2 *Instruction) int64 {
	switch i2 {
	case nil:
		return int64(i)
	}
	return int64(i)<<32 | i2.Int64(nil)
}

// Float64 returns the float32 representation of the instruction's bits
// or rather, it's immediate value
func (i Instruction) Float64(i2 Instruction) float32 {
	return *(*float32)(unsafe.Pointer(&i))
}
