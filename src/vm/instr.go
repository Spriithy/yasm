package vm

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
		println(opName[op], opName[rd], opName[rs0], int32(arg))
		return &iInstr{op, rd, rs0, int32(arg)}
	case op < _b: // R. Instructions
		rd := byte((src & rdMask) >> 21)
		rs0 := byte((src & rs0Mask) >> 16)
		rs1 := byte((src & rs1Mask) >> 11)
		_, _, b0, b1 := uint32ToBytes(src & rArgMask)
		println(opName[op], opName[rd], opName[rs0], opName[rs1], int16FromBytes(b0, b1))
		return &rInstr{op, rd, rs0, rs1, int32FromBytes(0, 0, b0, b1)}
	case op < opcodeMax: // B. Instructions
		_, b0, b1, b2 := uint32ToBytes(src & bArgMask)
		println(opName[op], int32FromBytes(0, b0, b1, b2))
		return &bInstr{op, int32FromBytes(0, b0, b1, b2)}
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
