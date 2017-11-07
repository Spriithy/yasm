package yasm

import (
	"fmt"
	"unsafe"
)

func Main() {
	c := &cpu{}
	c.sp = addr(unsafe.Pointer(&c.stack))
	c.sp += addr(len(c.stack))
	c.bp = c.sp - StackAlign
	c.fp = c.sp - StackAlign

	c.pushf64(0.8)
	c.pushf64(-1.2)
	c.dumpStack()
	fmt.Printf("fp=0x%x | *fp=0x%x\n", c.fp, getu32(c.fp))
	fmt.Printf("bp=0x%x | *bp=0x%x\n", c.bp, getu32(c.bp))
	fmt.Printf("sp=0x%x | *sp=0x%x\n", c.sp, getu32(c.sp))
}
