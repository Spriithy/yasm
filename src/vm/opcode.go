package vm

const (
	// Begin I. typed instructions with mov
	Mov = byte(iota) // Move

	Addi  // Add immediate
	Addiu // Add immediate Unsigned
	Negi  // Negate a signed number
	Andi  // And immediate
	Ori   // Or immediate
	Xori  // Not immediate

	// Branching related
	Beq // Branch equal
	Bne // Branch not equal
	Blt // Branch less than
	Bgt // Branch greater than
	Ble // Branch less or equal
	Bge // Branch greater or equal

	// Load operations
	Lb   // Load byte
	Lw   // Load word
	Lwu  // Load Word unsigned
	Ldw  // Load DoubleWord
	Ldwu // Load DoubleWord unsigned
	Lqw  // Load QuadWord
	Lqwu // Load QuadWord unsigned

	// Begin R. typed instructions
	Add  // Add
	Addu // Add unsigned
	Sub  // Subtract
	Subu // Subtract unsigned
	Mul  // Multiply
	Mulu // Multiply unsigned
	Div  // Divide
	Divu // Divide unsigned
	Shl  // Shift Right
	Shr  // Shift Left

	And // And
	Or  // Or
	Xor // Xor
	Not // Two's complement

	// Begin B. typed instructions
	B  // Direct Branch
	Br // Branch Register

	// Aliases for op type bounds
	_i = Mov
	_r = Add
	_b = B

	// Last valid opcode
	opcodeMax = 1<<6 - 1

	//
)

var opName = map[byte]string{
	Mov: "mov",

	Addi:  "addi",
	Addiu: "addiu",
	Negi:  "negi",
	Andi:  "andi",
	Ori:   "ori",
	Xori:  "xori",

	// Branching related
	Beq: "beq",
	Bne: "bne",
	Blt: "blt",
	Bgt: "bgt",
	Ble: "ble",
	Bge: "bge",

	// Load operations
	Lb:   "lb",
	Lw:   "lw",
	Lwu:  "lwu",
	Ldw:  "ldw",
	Ldwu: "ldwu",
	Lqw:  "lqw",
	Lqwu: "lqwu",

	// Begin R. typed instructions
	Add:  "add",
	Addu: "addu",
	Sub:  "sub",
	Subu: "subu",
	Mul:  "mul",
	Mulu: "mulu",
	Div:  "div",
	Divu: "divu",
	Shl:  "shl",
	Shr:  "shr",

	And: "and",
	Or:  "or",
	Xor: "xor",
	Not: "not",

	// Begin B. typed instructions
	B:  "not",
	Br: "br",
}
