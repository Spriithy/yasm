package yasm

import (
	"fmt"
	"os"
)

type errno = int

const (
	StackUnderflow = iota + 1
	StackOverflow
	ArithmeticOverflow
	errnoMax
)

var errnoStr = map[errno]string{
	StackUnderflow:     "StackUnderflow",
	StackOverflow:      "StackOverflow",
	ArithmeticOverflow: "ArithmeticOverflow",
}

var errnoCode = map[errno]int{
	StackUnderflow:     1,
	StackOverflow:      1,
	ArithmeticOverflow: 1,
}

func (c *cpu) trap(n errno) {
	if n >= errnoMax || n < 0 {
		fmt.Printf("\x1b[31merror: unknown trap at 0x%x\x1b[0m\n", c.ta)
		c.trace()
		os.Exit(errnoCode[n])
	}

	fmt.Printf("\x1b[31merror: %s at 0x%x\x1b[0m\n", errnoStr[n], c.ta)
	c.trace()
	os.Exit(errnoCode[n])
}
