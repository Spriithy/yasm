package vm

import "fmt"

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
	rdMask   uint32 = 0x3E00000
	rs0Mask  uint32 = 0x1F0000
	rs1Mask  uint32 = 0xF800
	iArgMask uint32 = 0xFFFF
	rArgMask uint32 = 0x7FF
	bArgMask uint32 = 0x3FFFFFF
)

func decode(src uint32) instr {
	op := byte(src >> 26)
	switch {
	case op > opcodeMax:
		return nil
	case op < _r: // Immediate instruction case
		rd := byte((src & rdMask) >> 21)
		rs0 := byte((src & rs0Mask) >> 16)
		arg := src & iArgMask
		return &iInstr{op, rd, rs0, int32(arg)}
	case op < _b: // R. Instructions
		rd := byte((src & rdMask) >> 21)
		rs0 := byte((src & rs0Mask) >> 16)
		rs1 := byte((src & rs1Mask) >> 11)
		x := Word(src & rArgMask)
		return &rInstr{op, rd, rs0, rs1, int32(x.Int16(0))}
	case op < opcodeMax: // B. Instructions
		x := Word(src & bArgMask)
		b0, b1, b2 := x.Int8(0), x.Int8(1), x.Int8(2)
		y := int32(b0) | int32(b1)<<8 | int32(b2)<<16
		return &bInstr{op, y}
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
	string() string
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

func (i *bInstr) string() string {
	return fmt.Sprintf("%5s\t#(0x%X)", opName[i.op], i.arg)
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

func (i *rInstr) string() string {
	return fmt.Sprintf("%5s\tr%d, r%d, r%d, #(0x%X)", opName[i.op], i.rd, i.rs0, i.rs1, i.arg)
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

func (i *iInstr) string() string {
	return fmt.Sprintf("%5s\tr%d, r%d, #(0x%X)", opName[i.op], i.rd, i.rs0, i.arg)
}
