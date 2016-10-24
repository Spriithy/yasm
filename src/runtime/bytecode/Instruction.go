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
	BR_T:    {"BRT", 2},
	BR_F:    {"BRF", 2},

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
	NOT:     {"NOT", 0},
	AND:     {"AND", 0},
	OR:      {"OR", 0},
	XOR:     {"XOR", 0},

	BEQ:     {"BEQ", 0},
	BNEQ:    {"BNEQ", 0},
	BLT:     {"BLT", 0},
	BGT:     {"BGT", 0},
	BLEQ:    {"BLEQ", 0},
	BGEQ:    {"BGEQ", 0},
	BADD:    {"BADD", 0},
	BSUB:    {"BSUB", 0},
	BMUL:    {"BMUL", 0},
	BDIV:    {"BDIV", 0},
	BMOD:    {"BMOD", 0},
	BSHR:    {"BSHR", 0},
	BSHL:    {"BSHL", 0},
	BCOMPL1: {"BCOMPL1", 0},
	BCOMPL2: {"BCOMPL2", 0},
	BAND:    {"BAND", 0},
	BOR:     {"BOR", 0},
	BXOR:    {"BXOR", 0},

	IEQ:     {"IEQ", 0},
	INEQ:    {"INEQ", 0},
	ILT:     {"ILT", 0},
	IGT:     {"IGT", 0},
	ILEQ:    {"ILEQ", 0},
	IGEQ:    {"IGEQ", 0},
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
	IXOR:    {"IXOR", 0},

	REQ:     {"REQ", 0},
	RNEQ:    {"RNEQ", 0},
	RLT:     {"RLT", 0},
	RGT:     {"RGT", 0},
	RLEQ:    {"RLEQ", 0},
	RGEQ:    {"RGEQ", 0},
	RADD:    {"RADD", 0},
	RSUB:    {"RSUB", 0},
	RMUL:    {"RMUL", 0},
	RDIV:    {"RDIV", 0},

	BCONST_0:{"BCONST0", 0},
	BCONST_1:{"BCONST1", 0},
	BCONST_2:{"BCONST2", 0},
	BCONST_3:{"BCONST3", 0},
	BCONST_4:{"BCONST4", 0},
	BCONST_N:{"BCONSTN", 1},

	ICONST_0:{"ICONST0", 0},
	ICONST_1:{"ICONST1", 0},
	ICONST_2:{"ICONST2", 0},
	ICONST_N:{"ICONSTN", 1},

	RCONST_0:{"RCONST0", 0},
	RCONST_1:{"RCONST1", 0},
	RCONST_2:{"RCONST2", 0},
	RCONST_N:{"RCONSTN", 1},
	RCONST_E:{"RCONST_E", 0},
	RCONST_PI:{"RCONST_PI", 0},
}
