package obj

import "github.com/Spriithy/Polaroid/reg/runtime/code"

type local struct {
	varname string
	startpc int
	endpc   int
}

type constant struct {
	T     RType
	Str   RString
	Byte  RByte
	Int   RInt
	UInt  RUInt
	Float RFloat
}

// RFunction is the internal representation of a function
type RFunction struct {
	source string
	line   int
	nstats byte
	nargs  byte
	varg   byte
	regc   byte

	iLines []int
	code   []code.Bytecode

	locals  []local
	statics []string
	consts  []constant
	funs    []RFunction
}
