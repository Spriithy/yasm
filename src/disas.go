package yasm

import (
	"fmt"
	"unsafe"
)

func (c *cpu) disasCurrent() {
	pc, code := int(c.fn.pc), c.Module.Code
	if fName, isFunc := c.Module.FuncNames[uintptr(pc)]; isFunc {
		fmt.Printf("0x%08x ", &code[pc])
		fmt.Printf(".%s \n", fName)
	}

	fmt.Printf("0x%08x %-12s", &code[pc], OpcodeName[code[pc]])
	switch code[pc] {
	case Syscall, Const8:
		fmt.Printf("$0x%x", code[pc+1])
	case Const16:
		fmt.Printf("$0x%x", *(*u16)(unsafe.Pointer(&code[pc+1])))
	case Const32:
		fmt.Printf("$0x%x", *(*u32)(unsafe.Pointer(&code[pc+1])))
	case Const64:
		fmt.Printf("$0x%x", *(*u64)(unsafe.Pointer(&code[pc+1])))
	case Pushl8, Pushl8u, Pushg8, Pushg8u, Popl8, Popg8,
		Pushl16, Pushl16u, Pushg16, Pushg16u, Popl16, Popg16,
		Pushl32, Pushl32u, Pushg32, Pushg32u, Popl32, Popg32,
		Pushl64, Pushg64, Popl64, Popg64:
		fmt.Printf("$%d", code[pc+1])
	}
}

func (m *Module) Disas() {
	for i := 0; i < len(m.Code); i++ {
		op := m.Code[i]

		if fName, isFunc := m.FuncNames[uintptr(i)]; isFunc {
			fmt.Printf("0x%08x ", &m.Code[i])
			fmt.Printf(".%s \n", fName)
		}

		fmt.Printf("0x%08x     ", &m.Code[i])
		fmt.Printf("%-12s", OpcodeName[m.Code[i]])
		switch op {
		case Syscall, Const8:
			fmt.Printf("$0x%x", m.Code[i+1])
			i += int(SizeOf[U8])
		case Const16:
			fmt.Printf("$0x%x", *(*u16)(unsafe.Pointer(&m.Code[i+1])))
			i += int(SizeOf[U16])
		case Const32:
			fmt.Printf("$0x%x", *(*u32)(unsafe.Pointer(&m.Code[i+1])))
			i += int(SizeOf[U32])
		case Const64:
			fmt.Printf("$0x%x", *(*u64)(unsafe.Pointer(&m.Code[i+1])))
			i += int(SizeOf[U64])
		case Pushl8, Pushl8u, Pushg8, Pushg8u, Popl8, Popg8,
			Pushl16, Pushl16u, Pushg16, Pushg16u, Popl16, Popg16,
			Pushl32, Pushl32u, Pushg32, Pushg32u, Popl32, Popg32,
			Pushl64, Pushg64, Popl64, Popg64:
			fmt.Printf("$%d", m.Code[i+1])
			i++
		}
		fmt.Println()
	}
}
