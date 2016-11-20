package runtime

const (
	// Memory / Registers
	mov = byte(iota) // Move
	swp              // Swap
	ptr              // Pointer to

	// Arithmetic
	ineg // Unary minus
	fneg
	ipow // Exponentiation
	fpow
	cadd // Addition
	iadd
	fadd
	csub // Subtraction
	isub
	fsub
	cmul // Multiplication
	imul
	fmul
	cdiv // Division
	idiv
	fdiv
	crem // Remainder
	irem
	cinc // Increment
	iinc
	cdec // Decrement
	idec
	cshr // Shift Right
	ishr
	cshl // Shift Left
	ishl

	// Logic
	eq
	neq
	clt // Less Than
	ilt
	flt
	cgt // Greater than
	igt
	fgt
	cleq // Less Equal Than
	ileq
	fleq
	cgeq // Greater Equal Than
	igeq
	fgeq
	or // Or
	cor
	ior
	and // And
	cand
	iand
	xor // Xor
	cxor
	ixor
	not // Not
	inot

	// Branching
	b   // Branch
	bz  // Branch if Zero
	bnz // Branch if Not Zero
	brt // Branch True
	brf // Branch False
	brn // Branch Null
	blp // Branch Last Program (counter) == Return

	//
)
