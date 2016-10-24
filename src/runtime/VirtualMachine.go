package runtime

import (
	"github.com/Spriithy/Polaroid/src/builtins"
	"github.com/Spriithy/Polaroid/src/runtime/bytecode"
	"github.com/Spriithy/Polaroid/src/runtime/metadata"
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
			v1 = v.Pop()
			addr = int(src[v.ip])
			switch v1.(type) {
			case builtins.Boolean:
				if v1.(builtins.Boolean) {
					v.ip = addr
				}
			default:
				// TODO VM ERROR
				return
			}
		case code.BR_F:
			v1 = v.Pop()
			addr = int(src[v.ip])
			switch v1.(type) {
			case builtins.Boolean:
				if !v1.(builtins.Boolean) {
					v.ip = addr
				}
			default:
				// TODO VM ERROR
				return
			}
		case code.IF_T:
			v1 = v.Pop()
			switch v1.(type) {
			case builtins.Boolean:
				if v1.(builtins.Boolean) {
					v.ip = int(src[v.ip])
				} else {
					v.ip = int(src[v.ip + 1])
				}
			default:
				// TODO VM ERROR
				return
			}
		case code.IF_F:
			v1 = v.Pop()
			switch v1.(type) {
			case builtins.Boolean:
				if !v1.(builtins.Boolean) {
					v.ip = int(src[v.ip])
				} else {
					v.ip = int(src[v.ip + 1])
				}
			default:
				// TODO VM ERROR
				return
			}
		case code.IF_NULL:
			v1 = v.Pop()
			switch v1.(type) {
			case builtins.Null:
				v.ip = int(src[v.ip])
			default:
				v.ip = int(src[v.ip + 1])
				return
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

		case code.AND:
			v1 = v.Pop()
			switch v1.(type) {
			case builtins.Boolean:

			case builtins.Integer:
			case builtins.Byte:
			default:
				// TODO VM ERROR
				return
			}
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