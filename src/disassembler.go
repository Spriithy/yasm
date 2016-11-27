package small

import (
	"fmt"
	"os"

	"github.com/Spriithy/go-colors"
)

func regName(r int) string {
	switch {
	case r < 8:
		return fmt.Sprintf("r%d", r)
	case r == 8:
		return "%pc"
	case r == 9:
		return "%sp"
	case r == 10:
		return "%ic"
	case r == 11:
		return "%rx"
	}
	return "(?)"
}

func getHex(a interface{}) string {
	return fmt.Sprintf("%#x", a)
}

func opName(op byte) string {
	switch op {
	case ADD:
		return "add"
	case SUB:
		return "sub"
	case MUL:
		return "mul"
	case DIV:
		return "div"
	case REM:
		return "rem"
	case BSL:
		return "bsl"
	case BSR:
		return "bsr"
	case INC:
		return "inc"
	case DEC:
		return "dex"
	case AND:
		return "and"
	case IOR:
		return "ior"
	case XOR:
		return "xor"
	case JMP:
		return "jmp"
	case JZ:
		return "jz"
	case JNZ:
		return "jnz"
	case JEQ:
		return "jeq"
	case JNE:
		return "jne"
	case JLT:
		return "jlt"
	case JLE:
		return "jle"
	case JGT:
		return "jgt"
	case JGE:
		return "jge"
	}
	return "(?)"
}

func err(a ...interface{}) {
	println("disas:\t", colors.Red(colors.None, a...))
	os.Exit(1)
}

// Disasemble is used to pretty-print (or dissassemble) a binary
// file compiled for the SMALL virtual-machine.
func Disasemble(x []Instruction) {
	var ii Instruction
	extra := false

	code := x
	for i := 0; i < len(code); i++ {
		ii = code[i]
		switch ii.O() {
		case SWI:
			print("swi\t")
			switch ii.X() {
			case 0x0:
				println(regName(ii.RA()))
			case 0x1:
				println(getHex(ii.Int16()))
			default:
				err("SWI takes no immediate value.")
			}
		case MOV:
			print("mov\t", regName(ii.RA()))
			switch ii.X() {
			case 0x0:
				println(",", regName(ii.RB()))
				println("-->", ii.RB())
			case 0x1:
				println(",", getHex(ii.Int16()))
			case 0x02:
				println(",", code[i+1].Int32())
			case 0x03:
				println(",", code[i+1].Int64(&code[i+2]))
			}
		case MOM:
			print("mom\t[", regName(ii.RA()))
			switch ii.X() {
			case 0x0:
				println("],", "["+regName(ii.RB())+"]")
			}
		case LOA:
			print("loa\t", regName(ii.RA()))
			print(", [", regName(ii.RB()), "]\n")
		case STR:
			print("str\t[", regName(ii.RA()), "]")
			extra = true
		case ADD, SUB, MUL, DIV, REM, BSL, BSR, INC, DEC:
			if ii.T() {
				print("i")
			} else {
				print("f")
			}
			fallthrough
		case AND, IOR, XOR:
			print(opName(ii.O()) + "\t")
			print(regName(ii.RA()))
			extra = true
		case NOT:
			print("not\t", ii.RA(), "\n")
		case JMP, JZ, JNZ:
			println(opName(ii.O())+"\t", ii.E())
		case JEQ, JNE, JLT, JLE, JGT, JGE:
			print(opName(ii.O()), "\t")
			switch ii.X() {
			case 0x0:
				println(",", regName(ii.RA()))
			default:
				println(",", ii.Int16())
			}
			println(",", code[i+1].Int32())
			i++
		case SRL:
			println("srl\t"+regName(ii.RA())+",", regName(ii.RB()))
		case RET:
			println("ret\t"+regName(ii.RA())+",", regName(ii.RB()))
		default:
			println("(?)")
		}

		if extra {
			extra = false
			switch ii.X() {
			case 0x0:
				println(",", regName(ii.RB()))
			case 0x1:
				println(",", getHex(ii.Int16()))
			case 0x02:
				println(",", code[i+1].Int32())
				i++
			case 0x03:
				println(",", code[i+1].Int64(&code[i+2]))
				i += 2
			}
		}
	}
}
