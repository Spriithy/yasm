package yasm

import (
	"unsafe"
)

func Main() {
	c := &cpu{}
	c.sp = uintptr(unsafe.Pointer(&c.stack))
	c.sp += StackSize
	c.bp = c.sp - StackAlign
	c.fp = c.sp - StackAlign

	c.Module = &Module{
		Name:      "test.yasm",
		Funcs:     make(map[string]uintptr),
		FuncNames: make(map[uintptr]string),
		Code: []Opcode{
			Pushl64, 0x04,
			Pushl64, 0x0c,
			Popl64, 0x04,
			Popl64, 0x0c,
			Syscall, 0x80,
			Unreachable,
		},
	}

	c.Module.RegisterFunc("main", 0)

	c.exec([]string{
		"test.yasm",
		"--version",
	})

	c.stackDump(0)

	c.Module.Disas()
}
