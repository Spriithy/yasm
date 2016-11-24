package vm

const (
	zro = 0x00 // zro	a, -	ra := 0
	mov = 0x01 // mov	a, b	ra := rb
	mom = 0x02 // mom	a, b	[ra] := [rb]

	loa = 0x03 // loa	a, b	ra := [rb]
	str = 0x04 // str	a, b	[ra] := rb

	add = 0x10 // add	a, b	rx := ra + rb
	adu = 0x11 // adu	a, b	ra := ra + rb (unsigned bits)

	sub = 0x12 // sub	a, b	ra := ra - rb
	sbu = 0x13 // sbu	a, b	ra := ra - rb (unsigned bits)

	mul = 0x14 // mul	a, b	ra := ra * rb
	mlu = 0x15 // mlx	a, b	ra := ra * rb (unsigned bits)

	div = 0x16 // div	a, b	ra := ra / rb
	dvu = 0x17 // dvu	a, b	ra := ra // rb ; rx := ra % rb

	mod = 0x18 // mod	a, b	ra := ra % rb

	inc = 0x19 // inc	a  -	ra++
	dec = 0x1a // dec	a  -	ra--

	gbs = 0x1b // gbs	a, b	ra := gbs(rb) (greatest bit set)

	and = 0x1c // and	a, b	ra := ra & rb
	ior = 0x1d // ior	a, b	ra := ra | rb
	xor = 0x1e // xor	a, b	ra := ra ^ rb

	shl = 0x1f // shl	a, b	ra := ra << rb
	shr = 0x20 // shr	a, b	ra := ra >> rb

	rol = 0x21 // rol	a, b	?
	ror = 0x22 // ror	a, b	?

	eql = 0x23 // eql	a, b	rx := (a == b)
	neq = 0x24 // neq	a, b	rx := (a != b)

	cmp = 0x25 // cmp	a, b	rx := -1 if (a < b), 0 if (a == b), 1 if (a > b)

	lt  = 0x26 // lt	a, b	rx := (a < b)
	leq = 0x27 // leq	a, b	rx := (a <= b)

	gt  = 0x28 // gt	a, b	rx := (a > b)
	geq = 0x29 // geq	a, b	rx := (a >= b)

	jmp = 0x2a // jmp	ofs 	pc := pc + ofs
	jeq = 0x2b // jeq	ofs		rx == 0 ? jmp ofs : pc++
	jne = 0x2c // jne	ofs 	rx == 0 ? pc++ : jmp ofs
	jlt = 0x2d // jlt	ofs 	rx <  0 ? jmp ofs : pc++
	jle = 0x2e // jle	ofs 	rx <= 0 ? jmp ofs : pc++
	jgt = 0x2f // jgt	ofs 	rx >  0 ? jmp ofs : pc++
	jge = 0x30 // jge	ofs 	rw >= 0 ? jmp ofs : pc++
)

var opName = map[byte]string{
	zro: "zro", mov: "mov", mom: "mom", loa: "loa", str: "str",
	add: "add", adu: "adu", sub: "sub", sbu: "sbu", mul: "mul",
	mlu: "mlu", div: "div", dvu: "dvu", mod: "mod", inc: "inc",
	dec: "dec", gbs: "gbs", and: "and", ior: "ior", xor: "xor",
	shl: "shl", shr: "shr", rol: "rol", ror: "ror", eql: "eql",
	neq: "neq", cmp: "cmp", lt: " lt", leq: "leq", gt: " gt",
	jmp: "jmp", jeq: "jeq", jne: "jne", jlt: "jlt", jle: "jle",
	jgt: "jgt", jge: "jge",
}
