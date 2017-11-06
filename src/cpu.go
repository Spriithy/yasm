package small

import (
	"fmt"
	"unsafe"
)

type cpu struct {
	stack [1024]u8
	sp    addr
	fp    addr
	bp    addr
	slp   addr // stack limit pointer

	ta addr // trap address
	ec int  // exit code

	text []Opcode
	data []u8
	dp   addr // data pointer
}

func (c *cpu) dumpMem() {
	for i := len(c.stack); i > 0; i-- {
		if c.sp-1 >= addr(unsafe.Pointer(&c.stack[i-1]))-stackAlign && c.sp-1 < addr(unsafe.Pointer(&c.stack[i-1])) {
			print("\x1b[1m")
		} else {
			print("\x1b[0m")
		}

		if i%int(stackAlign) == 0 {
			if i != len(c.stack) {
				print("\n")
			}
			fmt.Printf("0x%08x   ", i)
		}
		fmt.Printf(" 0x%02x", c.stack[i-1])
	}
	print("\n")
}
