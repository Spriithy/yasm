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

	c.module = &Module{
		Name:         "test.yasm",
		Libs:         nil,
		Funcs:        make(map[string]uintptr),
		Instructions: []Opcode{},
	}

	c.module.RegisterFunc("main", 0)

	c.exec([]string{
		"test.yasm",
		"version",
	})

	c.stackDump()
}
