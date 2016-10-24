package code

type Instruction struct {
	Name  string
	Nargs int
}

var InstructionTable = map[Bytecode]Instruction{
	NOOP:    {"NOP", 0},
	HALT:    {"HALT", 0},
	DROP:    {"DROP", 0},
	DUP:     {"DUP", 0},
	DUP2:    {"DUP2", 0},
	SWAP:    {"SWAP", 0},
	SWAP2:   {"SWAP2", 0},

	BR:      {"BR", 1},
	BR_T:    {"BRT", 1},
	BR_F:    {"BRF", 1},

	IF_T:    {"IFT", 2},
	IF_F:    {"IFF", 2},
	IF_NULL: {"IFNULL", 2},

	CALL:    {"CALL", 0},
	RETURN:  {"RET", 0},

	NEW:     {"NEW", 0},

	FGET:    {"FGET", 1},
	FSET:    {"FSET", 1},

	AGET:    {"AGET", 0},
	ASET:    {"ASET", 0},

	LOAD:    {"LOAD", 1},
	STORE:   {"STORE", 1},

	B2I:     {"B2I", 0},
	I2B:     {"I2B", 0},
	I2R:     {"I2R", 0},
	R2I:     {"R2I", 0},

	EQ:      {"EQ", 0},
	NEQ:     {"NEQ", 0},
	CMP:     {"CMP", 0},
	LT:      {"LT", 0},
	GT:      {"GT", 0},
	LEQ:     {"LEQ", 0},
	GEQ:     {"GEQ", 0},

	NOT:     {"NOT", 0},
	AND:     {"AND", 0},
	OR:      {"OR", 0},
	XOR:     {"XOR", 0},

	IADD:    {"IADD", 0},
	ISUB:    {"ISUB", 0},
	IMUL:    {"IMUL", 0},
	IDIV:    {"IDIV", 0},
	IMOD:    {"IMOD", 0},
	ISHR:    {"ISHR", 0},
	ISHL:    {"ISHL", 0},
	ICOMPL1: {"ICOMPL1", 0},
	ICOMPL2: {"ICOMPL2", 0},
	IAND:    {"IAND", 0},
	IOR:     {"IOR", 0},

	RADD:    {"RADD", 0},
	RSUB:    {"RSUB", 0},
	RMUL:    {"RMUL", 0},
	RDIV:    {"RDIV", 0},

	ICONST_0:{"ICONST0", 0},
	ICONST_1:{"ICONST1", 0},
	ICONST_2:{"ICONST2", 0},
	ICONST_N:{"ICONSTN", 0},
}
