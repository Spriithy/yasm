package vm

var ()

// Main is
func Main() {
	/*
		var x uint32 = uint32(Mov)<<26 | uint32(Blt)<<21 | uint32(Lb)<<16 | 1673
		y := decode(x)
		println(y.string())
		var r0, r1 *reg
		r0, r1 = new(reg), new(reg)
		*r0 = 10

		*r1 = 1009

		println(*r0, *r1, eq(r0, r1)) // false
		mov(r0, r1)
		println(*r0, *r1, eq(r0, r1)) // true
		*r1++
		println(*r0, *r1, eq(r0, r1)) //false
	*/

	x := gen(0x03, Add, 1, 49, 17, 0xff)
	print(read(x))

}
