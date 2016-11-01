package code

type Bytecode uint

const (
	//
	// Technical Opcdes
	//

	NOOP Bytecode = iota

	HALT

	DROP
	DUP
	DUP2
	SWAP
	SWAP2

	BR // Branch
	BR_T
	BR_F

	CALL
	RETURN

	LOAD
	STORE

	NEW

	FGET // Structure get
	FSET // Structure set

	AGET // Array Get
	ASET // Array Set

	B2I // Byte to Int
	B2L // Byte to Long
	I2B // Int to Byte
	I2L // Int to Long
	L2B // Long to Byte
	L2I // Long to Int
	L2R // Long to Real
	I2R // Int to Real
	R2I // Real to Int
	R2L // Real to Long

	//
	// Arithmetic & Logic
	//

	EQ
	NEQ

	NOT
	AND
	OR
	XOR

	BEQ
	BNEQ
	BLT
	BGT
	BLEQ
	BGEQ
	BADD
	BSUB
	BMUL
	BDIV
	BMOD
	BSHR
	BSHL
	BCOMPL1
	BCOMPL2
	BAND
	BOR
	BXOR

	IEQ
	INEQ
	ILT
	IGT
	ILEQ
	IGEQ
	IADD
	ISUB
	IMUL
	IDIV
	IMOD
	ISHR
	ISHL
	ICOMPL1
	ICOMPL2
	IAND
	IOR
	IXOR

	REQ
	RNEQ
	RLT
	RGT
	RLEQ
	RGEQ
	RADD
	RSUB
	RMUL
	RDIV

	ICONST_0
	ICONST_1
	ICONST_2
	ICONST_N

	BCONST_0
	BCONST_1
	BCONST_2
	BCONST_3
	BCONST_4
	BCONST_N

	RCONST_0
	RCONST_1
	RCONST_2
	RCONST_N
	RCONST_E
	RCONST_PI
)
