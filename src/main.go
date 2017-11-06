package yasm

import (
	"unsafe"
)

func Main() {
	c := &cpu{}
	c.sp = addr(unsafe.Pointer(&c.stack))
	c.slp = c.sp
	c.sp += addr(len(c.stack))
	c.bp = c.sp
	c.fp = c.sp

	c.pushf64(0.8)
	c.pushf64(1.2)
	c.pushf64(c.popf64() + c.popf64())
	c.pushu8(0xff)
	c.dumpMem()

	println(c.popf64())
}
