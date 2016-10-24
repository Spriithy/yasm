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

	IF_T
	IF_F
	IF_NULL

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
	I2B // Int to Byte

	I2R // Int to Real
	R2I // Real to Int

	//
	// Arithmetic & Logic
	//

	EQ
	NEQ
	CMP
	LT
	GT
	LEQ
	GEQ

	NOT
	AND
	OR
	XOR

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

	RADD
	RSUB
	RMUL
	RDIV

	ICONST_0
	ICONST_1
	ICONST_2
	ICONST_N
)