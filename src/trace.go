package yasm

import "fmt"

var StackTraceMaxDepth = 16

func (c *cpu) trace() {
	var depth int
	for fn := c.fn; fn.Caller != nil && depth < StackTraceMaxDepth; fn = fn.Caller {
		if depth == 0 {
			fmt.Printf("Stack trace: ")
		}
		fmt.Printf("in call to %s ~ \tPc=0x%x, Sp=0x%x, Fp=0x%x\n", fn.Name, fn.pc, c.sp, c.fp)
		depth++
	}

	if depth >= StackTraceMaxDepth {
		fmt.Printf("...\n")
	}
}
