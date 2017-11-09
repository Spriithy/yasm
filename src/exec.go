package yasm

import (
	"unsafe"
)

const (
	ExitSuccess = iota
	ExitFailure
)

func (c *cpu) exec(argv []string) int {
	// Check main exists
	if pc, ok := c.Module.Funcs["main"]; ok {
		c.fn = &Function{
			Name:   "main",
			Caller: nil,
			pc:     pc,
		}
	} else {
		c.SetTrapMessage("function 'main' could not be found in Module '" + c.Module.Name + "'")
		c.TrapAnonymous(NoEntryPoint)
	}

	// Push argc and pointers to each argv
	for narg := range argv {
		// *(fp + 4 + n*8) = argv[n]
		c.pushPtr(uintptr(unsafe.Pointer(&argv[len(argv)-narg-1])))
	}
	// *fp = argc
	c.push32(i32(len(argv)))

	// fp = sp
	c.fp = c.sp

	c.disasCurrent()
	ExecOpcode[c.Module.Code[c.fn.pc]](c)

	return ExitSuccess
}
