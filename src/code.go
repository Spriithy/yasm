package yasm

type Opcode = u8

const (
	Nop = iota
	Unreachable
	Syscall

	Const8
	Const16
	Const32
	Const64

	Load8
	Load8u
	Store8

	Load16
	Load16u
	Store16

	Load32
	Load32u
	Store32

	Load64
	Store64

	Drop
	Swap
	Dup

	Call
	Ret
	Fcall

	Not
	And
	Or
	Xor
	Shl
	Shr

	Incr
	Decr

	Add
	Sub
	Mul
	Div

	Addu
	Subu
	Mulu
	Divu

	Addf
	Subf
	Mulf
	Divf
)
