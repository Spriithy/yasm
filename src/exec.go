package yasm

import (
	"fmt"
	"unsafe"
)

const (
	ExitSuccess = iota
	ExitFailure
)

func (c *cpu) exec(argv []string) int {
	// Check main exists
	if pc, ok := c.module.Funcs["main"]; ok {
		c.fn = &Function{
			Name:   "main",
			Caller: nil,
			pc:     pc,
		}
	} else {
		c.SetTrapMessage("function 'main' could not be found in module '" + c.module.Name + "'")
		c.TrapAnonymous(NoEntryPoint)
	}

	// Push argc and pointers to each argv
	for narg := range argv {
		c.pushPtr(uintptr(unsafe.Pointer(&argv[len(argv)-narg-1])))
		fmt.Printf("argv[%d] : %s\n", len(argv)-narg-1, *(*string)(unsafe.Pointer(c.peekPtr())))
	}
	c.pushi32(i32(len(argv)))

	return ExitSuccess
}
