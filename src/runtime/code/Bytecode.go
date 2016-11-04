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
	L2F // Long to Real
	I2F // Int to Real
	F2I // Real to Int
	F2L // Real to Long

	//
	// Arithmetic & Logic
	//

	EQ
	NEQ

	LT
	LE
	GT
	GE

	NOT
	AND
	OR
	XOR

	ADD
	SUB
	MUL
	DIV
	REM
	SHR
	SHL
	CP1
	CP2
)
