package runtime

import (
	"fmt"
	"github.com/Spriithy/Polaroid/src/runtime/code"
	"log"
)

const (
	INITIAL_STACK_SIZE = 1 << 10
	STACK_GROW_SIZE = 1 << 6
	STACK_SHRINK_SIZE = 1 << 5
)

var stackExtensionChunk = make([]*Value, STACK_GROW_SIZE)

type vm struct {
	stack     []*Value
	sp        int

	code      []code.Bytecode
	ip        int

	frame     *ActivationRecord
	functions map[string]Function
}

func VirtualMachine() *vm {
	v := new(vm)
	v.stack = make([]*Value, INITIAL_STACK_SIZE)
	v.sp = -1
	v.ip = 0
	return v
}

func (v *vm) Exec(from int) {
	v.ip = from
	v.frame = Call(nil, v.functions["main"], -1)
	v.run()
}

func (v *vm) run() {
	var (
		op code.Bytecode
		v1, v2, v3, v4 *Value
		addr int
	)
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
			v.pop()
		case code.DUP:
			v1 = v.pop()
			v.push(v1); v.push(v1)
		case code.DUP2:
			v1 = v.pop(); v2 = v.pop()
			v.push(v2); v.push(v1)
			v.push(v2); v.push(v1)
		case code.SWAP:
			v1 = v.pop(); v2 = v.pop()
			v.push(v2); v.push(v1)
		case code.SWAP2:
			v1 = v.pop(); v2 = v.pop()
			v3 = v.pop(); v4 = v.pop()
			v.push(v2); v.push(v1); v.push(v4); v.push(v3)

		// Conditionals
		case code.BR:
			v.ip = (int)(src[v.ip])
		case code.BR_T:
			v1 := v.pop().Bool()
			addr = (int)(src[v.ip])
			if v1 {
				v.ip = addr
			}
		case code.BR_F:
			v1 := v.pop().Bool()
			addr = int(src[v.ip])
			if !v1 {
				v.ip = addr
			}

		case code.CALL:
		// TODO
		case code.RETURN:
			v.ip = v.frame.retip
			v.frame = v.frame.caller

		case code.LOAD: // TODO
		case code.STORE: // TODO

		case code.B2I:
			v1 := v.pop().Byte()
			v.push((int32)(v1))
		case code.B2L:
			v1 := v.pop().Byte()
			v.push((int64)(v1))
		case code.I2B:
			v1 := v.pop().Int()
			v.push((byte)(v1))
		case code.I2L:
			v1 := v.pop().Int()
			v.push((int64)(v1))
		case code.I2F:
			v1 := v.pop().Int()
			v.push((float64)(v1))
		case code.L2B:
			v1 := v.pop().Long()
			v.push((byte)(v1))
		case code.L2I:
			v1 := v.pop().Long()
			v.push((int32)(v1))
		case code.L2F:
			v1 := v.pop().Long()
			v.push((float64)(v1))
		case code.F2I:
			v1 := v.pop().Float()
			v.push((int32)(v1))
		case code.F2L:
			v1 := v.pop().Float()
			v.push((int64)(v1))

		case code.EQ:
			v.push(v.pop().Equals(v.pop()))
		case code.NEQ:
			v.push(!v.pop().Equals(v.pop()))
		case code.LT:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() < v2.Byte())
			case int32:
				v.push(v1.Int() < v2.Int())
			case int64:
				v.push(v1.Long() < v2.Long())
			case float64:
				v.push(v1.Float() < v2.Float())
			default:
				log.Panicf("Cannot compare %T and %T", v1.data, v2.data)
			}
		case code.LE:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() <= v2.Byte())
			case int32:
				v.push(v1.Int() <= v2.Int())
			case int64:
				v.push(v1.Long() <= v2.Long())
			case float64:
				v.push(v1.Float() <= v2.Float())
			default:
				log.Panicf("Cannot compare %T and %T", v1.data, v2.data)
			}
		case code.GT:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() > v2.Byte())
			case int32:
				v.push(v1.Int() > v2.Int())
			case int64:
				v.push(v1.Long() > v2.Long())
			case float64:
				v.push(v1.Float() > v2.Float())
			default:
				log.Panicf("Cannot compare %T and %T", v1.data, v2.data)
			}
		case code.GE:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() >= v2.Byte())
			case int32:
				v.push(v1.Int() >= v2.Int())
			case int64:
				v.push(v1.Long() >= v2.Long())
			case float64:
				v.push(v1.Float() >= v2.Float())
			default:
				log.Panicf("Cannot compare %T and %T", v1.data, v2.data)
			}

		case code.NOT:
			v1 = v.pop()
			switch v1.data.(type) {
			case byte, int32, int64, float64:
				v.push(v1.data == 0)
			case bool:
				v.push(!v1.Bool())
			default:
				log.Panicf("Cannot perform `not` on typed %T value", v1.data)
			}
		case code.AND:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() & v2.Byte())
			case int32:
				v.push(v1.Int() & v2.Int())
			case int64:
				v.push(v1.Long() & v2.Long())
			case bool:
				v.push(v1.Bool() && v2.Bool())
			default:
				log.Panicf("Cannot perform `and` on %T and %T", v1.data, v2.data)
			}
		case code.OR:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() | v2.Byte())
			case int32:
				v.push(v1.Int() | v2.Int())
			case int64:
				v.push(v1.Long() | v2.Long())
			case bool:
				v.push(v1.Bool() || v2.Bool())
			default:
				log.Panicf("Cannot perform `or` on %T and %T", v1.data, v2.data)
			}
		case code.XOR:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() ^ v2.Byte())
			case int32:
				v.push(v1.Int() ^ v2.Int())
			case int64:
				v.push(v1.Long() ^ v2.Long())
			case bool:
				v.push(v1.Bool() != v2.Bool())
			default:
				log.Panicf("Cannot perform `xor` on %T and %T", v1.data, v2.data)
			}

		case code.ADD:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() + v2.Byte())
			case int32:
				v.push(v1.Int() + v2.Int())
			case int64:
				v.push(v1.Long() + v2.Long())
			case float64:
				v.push(v1.Float() + v2.Float())
			case string:
				v.push(v1.String() + v2.String())
			default:
				log.Panicf("Cannot perform addition on %T and %T", v1.data, v2.data)
			}
		case code.SUB:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() - v2.Byte())
			case int32:
				v.push(v1.Int() - v2.Int())
			case int64:
				v.push(v1.Long() - v2.Long())
			case float64:
				v.push(v1.Float() - v2.Float())
			default:
				log.Panicf("Cannot perform subtraction on %T and %T", v1.data, v2.data)
			}
		case code.MUL:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() * v2.Byte())
			case int32:
				v.push(v1.Int() * v2.Int())
			case int64:
				v.push(v1.Long() * v2.Long())
			case float64:
				v.push(v1.Float() * v2.Float())
			default:
				log.Panicf("Cannot perform multiplication on %T and %T", v1.data, v2.data)
			}
		case code.DIV:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() / v2.Byte())
			case int32:
				v.push(v1.Int() / v2.Int())
			case int64:
				v.push(v1.Long() / v2.Long())
			case float64:
				v.push(v1.Float() / v2.Float())
			default:
				log.Panicf("Cannot perform division on %T and %T", v1.data, v2.data)
			}
		case code.REM:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() % v2.Byte())
			case int32:
				v.push(v1.Int() % v2.Int())
			case int64:
				v.push(v1.Long() % v2.Long())
			default:
				log.Panicf("Cannot perform subtraction on %T and %T", v1.data, v2.data)
			}
		case code.SHL:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() << v2.Byte())
			case int32:
				v.push(v1.Int() << (uint32)(v2.Int()))
			case int64:
				v.push(v1.Long() << (uint64)(v2.Long()))
			default:
				log.Panicf("Cannot perform left-shift on %T and %T", v1.data, v2.data)
			}
		case code.SHR:
			v2 = v.pop()
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(v1.Byte() >> v2.Byte())
			case int32:
				v.push(v1.Int() >> (uint32)(v2.Int()))
			case int64:
				v.push(v1.Long() >> (uint64)(v2.Long()))
			default:
				log.Panicf("Cannot perform right-shift on %T and %T", v1.data, v2.data)
			}
		case code.CP1:
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(^v1.Byte())
			case int32:
				v.push(^v1.Int())
			case int64:
				v.push(^v1.Long())
			default:
				log.Panicf("Cannot perform one's-complement on %T", v1.data)
			}
		case code.CP2:
			v1 = v.pop()
			switch v1.data.(type) {
			case byte:
				v.push(^v1.Byte() + 1)
			case int32:
				v.push(^v1.Int() + 1)
			case int64:
				v.push(^v1.Long() + 1)
			default:
				log.Panicf("Cannot perform two's-complement on %T", v1.data)
			}
		}
	}
}

func (v *vm) pop() *Value {
	if len(v.stack) - v.sp > STACK_SHRINK_SIZE {
		v.stack = v.stack[:len(v.stack) - STACK_SHRINK_SIZE]
	}
	item := v.stack[v.sp]
	v.sp--
	return item
}

func (v *vm) push(item interface{}) {
	if len(v.stack) - v.sp < 1 {
		v.stack = append(v.stack, stackExtensionChunk ...)
	}
	v.sp++
	v.stack[v.sp] = Wrap(item)
}

func (v *vm) peek() *Value {
	return v.stack[v.sp]
}

func (v *vm) empty() bool {
	return v.sp < 0
}

func (v *vm) StackString() string {
	str := "["
	for i := v.sp - 1; i >= 0; i-- {
		str += fmt.Sprintf("%v", v.stack[i].ToString())
		if i != 0 {
			str += ", "
		}
	}
	return str + "]"
}