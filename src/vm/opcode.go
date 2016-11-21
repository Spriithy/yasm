package vm

const (
	// Begin I. typed instructions with mov
	mov = byte(iota) // Move

	addi  // Add immediate
	addiu // Add immediate Unsigned
	negi  // Negate a signed number
	andi  // And immediate
	ori   // Or immediate
	xori  // Not immediate

	// Branching related
	beq // Branch equal
	bne // Branch not equal
	blt // Branch less than
	bgt // Branch greater than
	ble // Branch less or equal
	bge // Branch greater or equal

	// Load operations
	lb  // Load byte
	lc  // Load char
	li  // Load int
	liu // Load int unsigned
	ll  // Load long
	llu // Load long unsigned

	// Begin R. typed instructions
	add  // Add
	addu // Add unsigned
	sub  // Subtract
	subu // Subtract unsigned
	mul  // Multiply
	mulu // Multiply unsigned
	div  // Divide
	divu // Divide unsigned
	shl  // Shift Right
	shr  // Shift Left

	and // And
	or  // Or
	xor // Xor
	not // Two's complement

	// Begin B. typed instructions
	b  // Direct Branch
	br // Branch Register

	// Aliases for op type bounds
	_i = mov
	_r = add
	_b = b

	// Last valid opcode
	opcodeMax = 2<<6 - 1

	//
)

var opName = map[byte]string{
	mov: "mov",

	addi:  "addi",
	addiu: "addiu",
	negi:  "negi",
	andi:  "andi",
	ori:   "ori",
	xori:  "xori",

	// Branching related
	beq: "beq",
	bne: "bne",
	blt: "blt",
	bgt: "bgt",
	ble: "ble",
	bge: "bge",

	// Load operations
	lb:  "lb",
	lc:  "lc",
	li:  "li",
	liu: "liu",
	ll:  "ll",
	llu: "llu",

	// Begin R. typed instructions
	add:  "add",
	addu: "addu",
	sub:  "sub",
	subu: "subu",
	mul:  "mul",
	mulu: "mulu",
	div:  "div",
	divu: "divu",
	shl:  "shl",
	shr:  "shr",

	and: "and",
	or:  "or",
	xor: "xor",
	not: "not",

	// Begin B. typed instructions
	b:  "not",
	br: "br",
}
