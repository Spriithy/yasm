package code

// A Bytecode is the byte's representation of an internal instruction
// of the VM.
type Bytecode byte

const (
	Move Bytecode = iota
	LoadConst
	LoadBool
	LoadNull
	GetStatic
	GetGlobal
	GetTable
	SetGlobal
	SetStatic
	SetTable
	NewTable
	Self

	JumpTrue
	JumpFalse
	Jump

	Call
	TailCall
	Return

	OrBool
	AndBool
	XorBool
	NotBool

	AddInt
	SubInt
	MulInt
	DivInt
	PowInt
	RemInt
	NegInt
	ShlInt
	ShrInt
	IncInt
	DecInt
	OrInt
	AndInt
	XorInt
	Cp2Int
	EqInt
	NeqInt
	LtInt
	GtInt
	LeqInt
	GeqInt

	AddUInt
	SubUInt
	MulUInt
	DivUInt
	PowUInt
	RemUInt
	ShlUInt
	ShrUInt
	IncUInt
	DecUInt
	OrUInt
	AndUInt
	XorUInt
	EqUInt
	NeqUInt
	LtUInt
	GtUInt
	LeqUInt
	GeqUInt

	AddByte
	SubByte
	MulByte
	DivByte
	RemByte
	ShlByte
	ShrByte
	IncByte
	DecByte
	OrByte
	AndByte
	XorByte
	EqByte
	NeqByte
	LtByte
	GtByte
	LeqByte
	GeqByte

	AddFloat
	SubFloat
	MulFloat
	DivFloat
	PowFloat
	NegFloat
	EqFloat
	NeqFloat
	LtFloat
	GtFloat
	LeqFloat
	GeqFloat

	lastOp
)

var opName = map[Bytecode]string{
	Move:      "move",
	LoadConst: "load_const",
	LoadBool:  "load_bool",
	LoadNull:  "load_null",
	GetStatic: "get_static",
	GetGlobal: "get_global",
	GetTable:  "get_table",
	SetGlobal: "set_global",
	SetStatic: "set_static",
	SetTable:  "set_table",
	NewTable:  "new_table",
	Self:      "self",
	JumpTrue:  "jmp_true",
	JumpFalse: "jmp_false",
	Jump:      "jmp",
	Call:      "call",
	TailCall:  "tail_call",
	Return:    "return",
	OrBool:    "or",
	AndBool:   "and",
	XorBool:   "xor",
	NotBool:   "not",
	AddInt:    "i_add",
	SubInt:    "i_sub",
	MulInt:    "i_mul",
	DivInt:    "i_div",
	PowInt:    "i_pow",
	RemInt:    "i_rem",
	NegInt:    "i_neg",
	ShlInt:    "i_shl",
	ShrInt:    "i_shr",
	IncInt:    "i_inc",
	DecInt:    "i_dec",
	OrInt:     "i_or",
	AndInt:    "i_and",
	XorInt:    "i_xor",
	Cp2Int:    "i_cp2",
	EqInt:     "i_eq",
	NeqInt:    "i_neq",
	LtInt:     "i_lt",
	GtInt:     "i_gt",
	LeqInt:    "i_leq",
	GeqInt:    "i_geq",
	AddUInt:   "u_add",
	SubUInt:   "u_sub",
	MulUInt:   "u_mul",
	DivUInt:   "u_div",
	PowUInt:   "u_pow",
	RemUInt:   "u_rem",
	ShlUInt:   "u_shl",
	ShrUInt:   "u_shr",
	IncUInt:   "u_inc",
	DecUInt:   "u_dec",
	OrUInt:    "u_or",
	AndUInt:   "u_and",
	XorUInt:   "u_xor",
	EqUInt:    "u_eq",
	NeqUInt:   "u_neq",
	LtUInt:    "u_lt",
	GtUInt:    "u_gt",
	LeqUInt:   "u_leq",
	GeqUInt:   "u_geq",
	AddByte:   "b_add",
	SubByte:   "b_sub",
	MulByte:   "b_mul",
	DivByte:   "b_div",
	RemByte:   "b_rem",
	ShlByte:   "b_shl",
	ShrByte:   "b_shr",
	IncByte:   "b_inc",
	DecByte:   "b_dec",
	OrByte:    "b_or",
	AndByte:   "b_and",
	XorByte:   "b_xor",
	EqByte:    "b_eq",
	NeqByte:   "b_neq",
	LtByte:    "b_lt",
	GtByte:    "b_gt",
	LeqByte:   "b_leq",
	GeqByte:   "b_geq",
	AddFloat:  "f_add",
	SubFloat:  "f_sub",
	MulFloat:  "f_mul",
	DivFloat:  "f_div",
	PowFloat:  "f_pow",
	NegFloat:  "f_neg",
	EqFloat:   "f_eq",
	NeqFloat:  "f_neq",
	LtFloat:   "f_lt",
	GtFloat:   "f_gt",
	LeqFloat:  "f_leq",
	GeqFloat:  "f_geq",
}

func (b *Bytecode) String() string {
	if *b >= lastOp {
		return "noop"
	}
	return opName[*b]
}
