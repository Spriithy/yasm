package runtime

import (
	"github.com/Spriithy/Polaroid/old/builtins"
	"github.com/Spriithy/Polaroid/old/runtime/bytecode"
	"github.com/Spriithy/Polaroid/old/runtime/metadata"
	"math"
	"go/constant"
)

const (
	DEFAULT_STACK_SIZE int = 1 << 10
	DEFAULT_CALL_STACK_SIZE int = 1 << 10

	NULL builtins.Null = builtins.NEW_NULL
	TRUE builtins.Boolean = builtins.Boolean(true)
	FALSE builtins.Boolean = builtins.Boolean(false)
)

type vm struct {
	Stack

	ip       int
	code     []code.Bytecode

	ctx      *metadata.Context
	metadata []metadata.FunctionMetaData

	trace    bool
}

func (v *vm) Exec(ip int) {
	constant.Imag()

	v.ip = ip
	v.ctx = metadata.CreateContext(nil, v.metadata[0], 0)
	v.run()
}

func (v *vm) run() {
	var op code.Bytecode
	var v1, v2, v3, v4 builtins.Object
	var addr int

	src := v.code

	for {
		v.ip++
		if v.ip > len(v.code) || v.code[v.ip] == code.HALT {
			return
		}
		op = src[v.ip]

		switch op {
		// Basics
		case code.NOOP:
		case code.HALT:
			return
		case code.DROP:
			v.Pop()
		case code.DUP:
			v1 = v.Pop()
			v.Push(v1, v1)
		case code.DUP2:
			v1 = v.Pop(); v2 = v.Pop()
			v.Push(v2, v1, v2, v1)
		case code.SWAP:
			v1 = v.Pop(); v2 = v.Pop()
			v.Push(v2, v1)
		case code.SWAP2:
			v1 = v.Pop(); v2 = v.Pop()
			v3 = v.Pop(); v4 = v.Pop()
			v.Push(v2, v1, v4, v3)

		// Conditionals
		case code.BR:
			v.ip = int(src[v.ip])
		case code.BR_T:
			v1 = v.PopBoolean()
			addr = int(src[v.ip])
			if v1.(builtins.Boolean) {
				v.ip = addr
			}
		case code.BR_F:
			v1 = v.PopBoolean()
			addr = int(src[v.ip])
			if !v1.(builtins.Boolean) {
				v.ip = addr
			}

		case code.CALL:
			findex := int(src[v.ip])
			nargs := v.metadata[findex].Nargs
			v.ctx = metadata.CreateContext(v.ctx, v.metadata[findex], v.ip)
			for i := 0; i < nargs; i++ {
				v.ctx.Locals[i] = v.Pop()
			}
			v.ip = v.metadata[findex].Addr
		case code.RETURN:
			v.ip = v.ctx.Rip
			v.ctx = v.ctx.Caller

		case code.LOAD: v.load()
		case code.STORE: v.store()

		case code.B2I:
			v1 := v.PopByte()
			v.Push(builtins.Integer(v1))
		case code.I2B:
			v1 := v.PopInteger()
			v.Push(builtins.Byte(v1))
		case code.I2R:
			v1 := v.PopInteger()
			v.Push(builtins.Real(v1))
		case code.R2I:
			v1 := v.PopReal()
			v.Push(builtins.Integer(v1))

		case code.EQ:
			v1 = v.Pop(); v2 = v.Pop()
			if v1.Hash() != v2.Hash() {
				v.Push(FALSE)
			} else {
				v.Push(TRUE)
			}
		case code.NEQ:
			v1 = v.Pop(); v2 = v.Pop()
			if v1.Hash() != v2.Hash() {
				v.Push(TRUE)
			} else {
				v.Push(FALSE)
			}
		case code.NOT:
			v1 := v.PopBoolean()
			v.Push(!v1)
		case code.OR:
			v1 := v.PopBoolean()
			v2 := v.PopBoolean()
			v.Push(v1 || v2)
		case code.AND:
			v1 := v.PopBoolean()
			v2 := v.PopBoolean()
			v.Push(v1 && v2)
		case code.XOR:
			v1 := v.PopBoolean()
			v2 := v.PopBoolean()
			v.Push(builtins.Boolean(v1 != v2))

		case code.BEQ:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 == v2))
		case code.BNEQ:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 != v2))
		case code.BLT:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 < v2))
		case code.BGT:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 > v2))
		case code.BLEQ:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 <= v2))
		case code.BGEQ:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(builtins.Boolean(v1 >= v2))
		case code.BADD:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 + v2)
		case code.BSUB:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 - v2)
		case code.BMUL:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 * v2)
		case code.BDIV:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 / v2)
		case code.BMOD:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 % v2)
		case code.BSHR:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 >> uint64(v2))
		case code.BSHL:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 << uint64(v2))
		case code.BCOMPL1:
			v1 := v.PopByte()
			v.Push(^v1)
		case code.BCOMPL2:
			v1 := v.PopByte()
			v.Push(^v1 + 1)
		case code.BAND:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 & v2)
		case code.BOR:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 | v2)
		case code.BXOR:
			v2 := v.PopByte()
			v1 := v.PopByte()
			v.Push(v1 ^ v2)

		case code.IEQ:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 == v2))
		case code.INEQ:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 != v2))
		case code.ILT:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 < v2))
		case code.IGT:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 > v2))
		case code.ILEQ:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 <= v2))
		case code.IGEQ:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(builtins.Boolean(v1 >= v2))
		case code.IADD:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 + v2)
		case code.ISUB:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 - v2)
		case code.IMUL:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 * v2)
		case code.IDIV:
			v1 := v.PopInteger()
			v2 := v.PopInteger()
			v.Push(v1 / v2)
		case code.IMOD:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 % v2)
		case code.ISHR:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 >> uint64(v2))
		case code.ISHL:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 << uint64(v2))
		case code.ICOMPL1:
			v1 := v.PopInteger()
			v.Push(^v1)
		case code.ICOMPL2:
			v1 := v.PopInteger()
			v.Push(^v1 + 1)
		case code.IAND:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 & v2)
		case code.IOR:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 | v2)
		case code.IXOR:
			v2 := v.PopInteger()
			v1 := v.PopInteger()
			v.Push(v1 ^ v2)

		case code.REQ:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 == v2))
		case code.RNEQ:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 != v2))
		case code.RLT:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 < v2))
		case code.RGT:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 > v2))
		case code.RLEQ:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 <= v2))
		case code.RGEQ:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(builtins.Boolean(v1 >= v2))
		case code.RADD:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(v1 + v2)
		case code.RSUB:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(v1 - v2)
		case code.RMUL:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(v1 * v2)
		case code.RDIV:
			v2 := v.PopReal()
			v1 := v.PopReal()
			v.Push(v1 / v2)

		case code.BCONST_0:  v.Push(builtins.Byte(0))
		case code.BCONST_1:  v.Push(builtins.Byte(1))
		case code.BCONST_2:  v.Push(builtins.Byte(2))
		case code.BCONST_3:  v.Push(builtins.Byte(3))
		case code.BCONST_4:  v.Push(builtins.Byte(4))
		case code.BCONST_N://v.Push(builtins.Byte(0))

		case code.ICONST_0:  v.Push(builtins.Integer(0))
		case code.ICONST_1:  v.Push(builtins.Integer(1))
		case code.ICONST_2:  v.Push(builtins.Integer(2))
		case code.ICONST_N://v.Push(builtins.Integer(0))

		case code.RCONST_0:  v.Push(builtins.Real(0.0))
		case code.RCONST_1:  v.Push(builtins.Real(1.0))
		case code.RCONST_2:  v.Push(builtins.Real(2.0))
		case code.RCONST_N://v.Push(builtins.Real(0.0))
		case code.RCONST_E:  v.Push(builtins.Real(math.E))
		case code.RCONST_PI: v.Push(builtins.Real(math.Pi))
		}
	}
}

func (v *vm) load() {
	regnum := v.code[v.ip]
	v.Push(v.ctx.Locals[regnum])
	v.ip++
}

func (v *vm) store() {
	regnum := v.code[v.ip]
	v.ctx.Locals[regnum] = v.Pop()
	v.ip++
}