package yasm

type Opcode = u8

const (
	nop = iota
	unreachable
	syscall

	const8
	const16
	const32
	const64

	load8
	load8u
	store8

	load16
	load16u
	store16

	load32
	load32u
	store32

	load64
	store64

	glocal
	slocal

	gglobal
	sglobal

	drop
	swap
	dup

	not
	and
	or
	xor
	shl
	shr

	incr
	decr

	add
	sub
	mul
	div

	addu
	subu
	mulu
	divu

	addf
	subf
	mulf
	divf
)
