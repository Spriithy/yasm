package yasm

import (
	"fmt"
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
		"--version",
	})

	c.fp = c.sp

	c.pushLocalPtr(4)
	c.pushLocalPtr(12)

	c.stackDump(0)

	fmt.Printf("argv[0] : %s\n", *(*string)(unsafe.Pointer(c.getLocalPtr(4))))
	fmt.Printf("argv[1] : %s\n", *(*string)(unsafe.Pointer(c.getLocalPtr(12))))

	c.popLocalPtr(4)
	c.popLocalPtr(12)
	fmt.Print("swap...")

	c.stackDump(0)

	fmt.Printf("argv[0] : %s\n", *(*string)(unsafe.Pointer(c.getLocalPtr(4))))
	fmt.Printf("argv[1] : %s\n", *(*string)(unsafe.Pointer(c.getLocalPtr(12))))
}
