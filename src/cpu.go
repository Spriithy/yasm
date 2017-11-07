package yasm

import (
	"fmt"
	"unsafe"
)

type cpu struct {
	stack [StackSize]u8
	sp    addr
	fp    addr
	bp    addr

	ta addr // trap address
	ec int  // exit code

	libs   map[string]*Module
	module *Module
	fn     *Function
}

func isDumpAt(p addr, pos addr) bool {
	if p-1 >= pos-StackAlign && p-1 < pos {
		return true
	}
	return false
}

func (c *cpu) dumpStack() {
	for i := len(c.stack); i > 1024-128; i-- {
		// dump pointer
		dp := addr(unsafe.Pointer(&c.stack[i-1]))
		if isDumpAt(c.sp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				print("\n")
			}
			fmt.Printf("\x1b[44;1m0x%08x SP \x1b[0m\x1b[34m\uE0B0\x1b[0m", dp-StackAlign+1)
		} else if isDumpAt(c.fp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				print("\n")
			}
			fmt.Printf("\x1b[42;1m0x%08x FP \x1b[0m\x1b[32m\uE0B0\x1b[0m", dp-StackAlign+1)
		} else if isDumpAt(c.bp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				print("\n")
			}
			fmt.Printf("\x1b[41;1m0x%08x BP \x1b[0m\x1b[31m\uE0B0\x1b[0m", dp-StackAlign+1)
		} else if i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				print("\n")
			}
			fmt.Printf("0x%08x     ", dp-StackAlign+1)
		}
		fmt.Printf(" 0x%02x", c.stack[i-1])
	}
	print("\n")
}
