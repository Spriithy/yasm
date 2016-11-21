package vm

import "strconv"

// BRI Instruction format
//
// B. Instructions
//
//    6 bits        26 bits
//    opcode    signed integer
//
// R. Instructions
//
//    6 bits     5 bits    5 bits  5 bits      11 bits
//    opcode      RD        RS0     RS1        extras
//
// I. Instructions
//
//    6 bits     5 bits    5 bits       16 bits
//    opcode      RD        RS0      signed integer
//

var (
	rdMask   int32 = 0x3E00000
	rs0Mask  int32 = 0x1F0000
	rs1Mask  int32 = 0xF800
	iArgMask int32 = 0xFFFF
	rArgMask int32 = 0x7FF
)

func decode(src int32) instr {
	op := byte(src >> 26)
	println(strconv.FormatInt(int64(src), 2))
	switch {
	case op > opcodeMax:
		return nil
	case op < _r: // Immediate instruction case
		rd := byte((src & rdMask) >> 21)
		rs0 := byte((src & rs0Mask) >> 16)
		arg := int32((src & iArgMask))
		println(opName[op], opName[rd], opName[rs0], arg)
		return &iInstr{op, rd, rs0, arg}
	case op < _b: // R. Instructions
		rd := byte((src & rdMask) >> 21)
		rs0 := byte((src & rs0Mask) >> 16)
		rs1 := byte((src & rs1Mask) >> 11)
		arg := src & rArgMask
		println(opName[op], opName[rd], opName[rs0], opName[rs1], arg)
		return &rInstr{op, rd, rs0, rs1, arg}
	case op < opcodeMax:
		arg := src << 6 >> 26
		println(opName[op], arg)
		return &bInstr{op, arg}
	}
	return nil
}

const (
	bFormat = byte(iota)
	rFormat
	iFormat
)

type instr interface {
	format() byte
	opcode() byte
}

type bInstr struct {
	op  byte
	arg int32
}

type rInstr struct {
	op  byte
	rd  byte
	rs0 byte
	rs1 byte
	arg int32
}

type iInstr struct {
	op  byte
	rd  byte
	rs0 byte
	arg int32
}

func (i *bInstr) format() byte {
	return bFormat
}

func (i *bInstr) opcode() byte {
	return i.op
}

func (i *bInstr) extra() int32 {
	return i.arg
}

func (i *rInstr) format() byte {
	return rFormat
}

func (i *rInstr) opcode() byte {
	return i.op
}

func (i *rInstr) dest() byte {
	return i.rd
}

func (i *rInstr) src0() byte {
	return i.rs0
}

func (i *rInstr) src1() byte {
	return i.rs1
}

func (i *iInstr) format() byte {
	return iFormat
}

func (i *iInstr) opcode() byte {
	return i.op
}

func (i *iInstr) dest() byte {
	return i.rd
}

func (i *iInstr) src0() byte {
	return i.rs0
}

func (i *iInstr) extra() int32 {
	return i.arg
}
