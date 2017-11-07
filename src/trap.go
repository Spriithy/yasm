package yasm

import (
	"fmt"
	"os"
)

type errno = int

const (
	StackUnderflow = iota + 1
	StackOverflow
	NoEntryPoint
	ArithmeticOverflow
	errnoMax
)

var errnoStr = map[errno]string{
	StackUnderflow:     "StackUnderflow",
	StackOverflow:      "StackOverflow",
	NoEntryPoint:       "NoEntryPoint",
	ArithmeticOverflow: "ArithmeticOverflow",
}

var errnoCode = map[errno]int{
	StackUnderflow:     ExitFailure,
	StackOverflow:      ExitFailure,
	NoEntryPoint:       ExitFailure,
	ArithmeticOverflow: ExitFailure,
}

func (c *cpu) SetTrapMessage(message string) {
	c.tm = message
}

func (c *cpu) Trap(n errno) {
	if n >= errnoMax || n < 0 {
		fmt.Printf("\x1b[31merror: unknown Trap at 0x%x\x1b[0m\n", c.ta)
		if c.tm != "" {
			fmt.Printf("  ~ %s\n", c.tm)
		}
		c.trace()
		os.Exit(errnoCode[n])
	}

	fmt.Printf("\x1b[31merror: %s at 0x%x\x1b[0m\n", errnoStr[n], c.ta)
	if c.tm != "" {
		fmt.Printf("  ~ %s\n", c.tm)
	}
	c.trace()
	os.Exit(errnoCode[n])
}

func (c *cpu) TrapAnonymous(n errno) {
	if n >= errnoMax || n < 0 {
		fmt.Printf("\x1b[31merror: unknown Trap at 0x%x\x1b[0m\n", c.ta)
		if c.tm != "" {
			fmt.Printf("  ~ %s\n", c.tm)
		}
		os.Exit(errnoCode[n])
	}

	fmt.Printf("\x1b[31merror: %s at 0x%x\x1b[0m\n", errnoStr[n], c.ta)
	if c.tm != "" {
		fmt.Printf("  ~ %s\n", c.tm)
	}
	os.Exit(errnoCode[n])
}
