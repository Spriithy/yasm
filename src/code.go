package yasm

type Opcode = u8

const (
	Nop = iota
	Unreachable
	Syscall

	Const8
	Const16
	Const32
	Const64

	Load8
	Load8u
	Store8

	Load16
	Load16u
	Store16

	Load32
	Load32u
	Store32

	Load64
	Store64

	Pushl8
	Pushl8u
	Pushg8
	Pushg8u
	Popl8
	Popg8

	Pushl16
	Pushl16u
	Pushg16
	Pushg16u
	Popl16
	Popg16

	Pushl32
	Pushl32u
	Pushg32
	Pushg32u
	Popl32
	Popg32

	Pushl64
	Pushg64
	Popl64
	Popg64

	Drop
	Drop64
	Swap
	Swap64
	Dup
	Dup64

	Call
	Ret

	Not
	And
	Or
	Xor
	Shl
	Shr
	Add
	Sub
	Mul
	Div
	Addu
	Subu
	Mulu
	Divu
	Addf
	Subf
	Mulf
	Divf

	Not64
	And64
	Or64
	Xor64
	Shl64
	Shr64
	Add64
	Sub64
	Mul64
	Div64
	Add64u
	Sub64u
	Mul64u
	Div64u
	Add64f
	Sub64f
	Mul64f
	Div64f
	opcodeMax
)

var OpcodeName = map[Opcode]string{
	Nop:         "nop",
	Unreachable: "unreachable",
	Syscall:     "syscall",
	Const8:      "u8.const",
	Const16:     "u16.const",
	Const32:     "u32.const",
	Const64:     "u64.const",
	Load8:       "i8.load",
	Load8u:      "u8.load",
	Store8:      "u8.store",
	Load16:      "i16.load",
	Load16u:     "u16.load",
	Store16:     "u16.store",
	Load32:      "i32.load",
	Load32u:     "u32.load",
	Store32:     "u32.store",
	Load64:      "u64.load",
	Store64:     "u64.store",
	Pushl8:      "i8.lload",
	Pushl8u:     "u8.lload",
	Pushg8:      "i8.gload",
	Pushg8u:     "u8.gload",
	Popl8:       "u8.lstore",
	Popg8:       "u8.gstore",
	Pushl16:     "i16.lload",
	Pushl16u:    "u16.lload",
	Pushg16:     "i16.gload",
	Pushg16u:    "u16.gload",
	Popl16:      "u16.lstore",
	Popg16:      "u16.gstore",
	Pushl32:     "i32.lload",
	Pushl32u:    "u32.lload",
	Pushg32:     "i32.gload",
	Pushg32u:    "u32.gload",
	Popl32:      "u32.lstore",
	Popg32:      "u32.gstore",
	Pushl64:     "u64.lload",
	Pushg64:     "u64.gload",
	Popl64:      "u64.lstore",
	Popg64:      "u64.gstore",
	Drop:        "u32.drop",
	Drop64:      "u64.drop",
	Swap:        "u32.swap",
	Swap64:      "u64.swap",
	Dup:         "u32.dup",
	Dup64:       "u64.dup",
	Call:        "call",
	Ret:         "ret",
	Not:         "u32.not",
	And:         "u32.and",
	Or:          "u32.or",
	Xor:         "u32.xor",
	Shl:         "u32.shl",
	Shr:         "u32.shr",
	Add:         "i32.add",
	Sub:         "i32.sub",
	Mul:         "i32.mul",
	Div:         "i32.div",
	Addu:        "u32.add",
	Subu:        "u32.sub",
	Mulu:        "u32.mul",
	Divu:        "u32.div",
	Addf:        "f32.add",
	Subf:        "f32.sub",
	Mulf:        "f32.mul",
	Divf:        "f32.div",
	Not64:       "u64.not",
	And64:       "u64.and",
	Or64:        "u64.or",
	Xor64:       "u64.xor",
	Shl64:       "u64.shl",
	Shr64:       "u64.shr",
	Add64:       "i64.add",
	Sub64:       "i64.sub",
	Mul64:       "i64.mul",
	Div64:       "i64.div",
	Add64u:      "u64.add",
	Sub64u:      "u64.sub",
	Mul64u:      "u64.mul",
	Div64u:      "u64.div",
	Add64f:      "f64.add",
	Sub64f:      "f64.sub",
	Mul64f:      "f64.mul",
	Div64f:      "f64.div",
}

var ExecOpcode = [opcodeMax]func(c *cpu){
	Nop: func(c *cpu) {},
	Unreachable: func(c *cpu) {
		c.Trap(UnreachableError)
	},
	Syscall: func(c *cpu) {
		c.SetTrapMessage("syscall not implemented")
		c.Trap(UnknownOpcodeError)
	},
	Const8: func(c *cpu) {
		c.push8u(c.Module.Code[c.fn.pc+1])
		c.fn.pc += SizeOf[U8]
	},
	Const16: func(c *cpu) {
		c.push16u(deref16u(c.Module.Code, c.fn.pc+1))
		c.fn.pc += SizeOf[U16]
	},
	Const32: func(c *cpu) {
		c.push32u(deref32u(c.Module.Code, c.fn.pc+1))
		c.fn.pc += SizeOf[U32]
	},
	Const64: func(c *cpu) {
		c.push64u(deref64u(c.Module.Code, c.fn.pc+1))
		c.fn.pc += SizeOf[U64]
	},
	Load8: func(c *cpu) {
		addr := c.popPtr()
		c.push8(get8(addr))
	},
	Load8u: func(c *cpu) {
		addr := c.popPtr()
		c.push8u(get8u(addr))
	},
	Store8: func(c *cpu) {
		set8u(c.popPtr(), c.pop8u())
	},
	Load16: func(c *cpu) {
		addr := c.popPtr()
		c.push16(get16(addr))
	},
	Load16u: func(c *cpu) {
		addr := c.popPtr()
		c.push16u(get16u(addr))
	},
	Store16: func(c *cpu) {
		set16u(c.popPtr(), c.pop16u())
	},
	Load32: func(c *cpu) {
		addr := c.popPtr()
		c.push32(get32(addr))
	},
	Load32u: func(c *cpu) {
		addr := c.popPtr()
		c.push32u(get32u(addr))
	},
	Store32: func(c *cpu) {
		set32u(c.popPtr(), c.pop32u())
	},
	Load64: func(c *cpu) {
		addr := c.popPtr()
		c.push64(get64(addr))
	},
	Store64: func(c *cpu) {
		set64u(c.popPtr(), c.pop64u())
	},
	Pushl8: func(c *cpu) {
		c.pushLocal8(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl8u: func(c *cpu) {
		c.pushLocal8u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg8: func(c *cpu) {
		c.pushGlobal8(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg8u: func(c *cpu) {
		c.pushGlobal8u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popl8: func(c *cpu) {
		c.popLocal8(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popg8: func(c *cpu) {
		c.popLocal8u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl16: func(c *cpu) {
		c.pushLocal16(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl16u: func(c *cpu) {
		c.pushLocal16u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg16: func(c *cpu) {
		c.pushGlobal16(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg16u: func(c *cpu) {
		c.pushGlobal16u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popl16: func(c *cpu) {
		c.popLocal16(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popg16: func(c *cpu) {
		c.popGlobal16(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl32: func(c *cpu) {
		c.pushLocal32(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl32u: func(c *cpu) {
		c.pushLocal32u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg32: func(c *cpu) {
		c.pushGlobal32(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg32u: func(c *cpu) {
		c.pushGlobal32u(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popl32: func(c *cpu) {
		c.popLocal32(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popg32: func(c *cpu) {
		c.popGlobal32(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushl64: func(c *cpu) {
		c.pushLocal64(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Pushg64: func(c *cpu) {
		c.pushGlobal64(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popl64: func(c *cpu) {
		c.popLocal64(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Popg64: func(c *cpu) {
		c.popGlobal64(int(c.Module.Code[c.fn.pc+1]))
		c.fn.pc++
	},
	Drop: func(c *cpu) {
		c.pop32()
	},
	Drop64: func(c *cpu) {
		c.pop64()
	},
	Swap: func(c *cpu) {
		x, y := c.pop32(), c.pop32()
		c.push32(y)
		c.push32(x)
	},
	Swap64: func(c *cpu) {
		x, y := c.pop64(), c.pop64()
		c.push64(y)
		c.push64(x)
	},
	Dup: func(c *cpu) {
		v := c.pop32()
		c.push32(v)
		c.push32(v)
	},
	Dup64: func(c *cpu) {
		v := c.pop64()
		c.push64(v)
		c.push64(v)
	},
	Call: func(c *cpu) {
		c.pushPtr(c.fn.pc + 4)

	},
	Ret: func(c *cpu) {
		// TODO
		c.SetTrapMessage("ret not implemented")
		c.Trap(UnknownOpcodeError)
	},
	Not: func(c *cpu) {
		x := c.pop32u()
		c.push32u(-x - 1) // C equivalent to ~x
	},
	And: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x & y)
	},
	Or: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x | y)
	},
	Xor: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x ^ y)
	},
	Shl: func(c *cpu) {
		x, n := c.pop32u(), c.Module.Code[c.fn.pc+1]
		c.push32u(x << n)
		c.fn.pc++
	},
	Shr: func(c *cpu) {
		x, n := c.pop32u(), c.Module.Code[c.fn.pc+1]
		c.push32u(x >> n)
		c.fn.pc++
	},
	Add: func(c *cpu) {
		x, y := c.pop32(), c.pop32()
		c.push32(x + y)
	},
	Sub: func(c *cpu) {
		x, y := c.pop32(), c.pop32()
		c.push32(x - y)
	},
	Mul: func(c *cpu) {
		x, y := c.pop32(), c.pop32()
		c.push32(x * y)
	},
	Div: func(c *cpu) {
		x, y := c.pop32(), c.pop32()
		c.push32(x / y)
	},
	Addu: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x + y)
	},
	Subu: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x - y)
	},
	Mulu: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x * y)
	},
	Divu: func(c *cpu) {
		x, y := c.pop32u(), c.pop32u()
		c.push32u(x / y)
	},
	Addf: func(c *cpu) {
		x, y := c.pop32f(), c.pop32f()
		c.push32f(x + y)
	},
	Subf: func(c *cpu) {
		x, y := c.pop32f(), c.pop32f()
		c.push32f(x - y)
	},
	Mulf: func(c *cpu) {
		x, y := c.pop32f(), c.pop32f()
		c.push32f(x * y)
	},
	Divf: func(c *cpu) {
		x, y := c.pop32f(), c.pop32f()
		c.push32f(x / y)
	},
	Not64: func(c *cpu) {
		x := c.pop64u()
		c.push64u(-x - 1) // C equivalent to ~x
	},
	And64: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x & y)
	},
	Or64: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x | y)
	},
	Xor64: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x ^ y)
	},
	Shl64: func(c *cpu) {
		x, n := c.pop64u(), c.Module.Code[c.fn.pc+1]
		c.push64u(x << n)
		c.fn.pc++
	},
	Shr64: func(c *cpu) {
		x, n := c.pop64u(), c.Module.Code[c.fn.pc+1]
		c.push64u(x >> n)
		c.fn.pc++
	},
	Add64: func(c *cpu) {
		x, y := c.pop64(), c.pop64()
		c.push64(x + y)
	},
	Sub64: func(c *cpu) {
		x, y := c.pop64(), c.pop64()
		c.push64(x - y)
	},
	Mul64: func(c *cpu) {
		x, y := c.pop64(), c.pop64()
		c.push64(x * y)
	},
	Div64: func(c *cpu) {
		x, y := c.pop64(), c.pop64()
		c.push64(x / y)
	},
	Add64u: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x + y)
	},
	Sub64u: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x - y)
	},
	Mul64u: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x * y)
	},
	Div64u: func(c *cpu) {
		x, y := c.pop64u(), c.pop64u()
		c.push64u(x / y)
	},
	Add64f: func(c *cpu) {
		x, y := c.pop64f(), c.pop64f()
		c.push64f(x + y)
	},
	Sub64f: func(c *cpu) {
		x, y := c.pop64f(), c.pop64f()
		c.push64f(x - y)
	},
	Mul64f: func(c *cpu) {
		x, y := c.pop64f(), c.pop64f()
		c.push64f(x * y)
	},
	Div64f: func(c *cpu) {
		x, y := c.pop64f(), c.pop64f()
		c.push64f(x / y)
	},
}
