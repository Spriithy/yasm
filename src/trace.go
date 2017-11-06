package yasm

import "fmt"

func (c *cpu) trace() {
	if c.fp == c.bp {
		// frame is first _start
		fmt.Printf("in _start:    sp=0x%x fp=0x%x\n", c.sp, c.fp)
		return
	}

	fmt.Printf("in %s:    sp=0x%x fp=0x%x\n", "<...>", c.sp, c.fp)
	c.fp = getAddr(c.fp + sizeOf[Addr])
	c.trace()
}
