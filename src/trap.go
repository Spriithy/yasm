package yasm

import (
	"fmt"
	"os"
)

const (
	StackUnderflow = iota + 1
	StackOverflow
	UnreachableError
	UnknownOpcodeError
	NoEntryPoint
	ArithmeticOverflow
	errnoMax
)

var errnoStr = map[int]string{
	StackUnderflow:     "StackUnderflow",
	StackOverflow:      "StackOverflow",
	UnreachableError:   "UnreachableError",
	UnknownOpcodeError: "UnknownOpcodeError",
	NoEntryPoint:       "NoEntryPoint",
	ArithmeticOverflow: "ArithmeticOverflow",
}

var errnoCode = map[int]int{
	StackUnderflow:     ExitFailure,
	StackOverflow:      ExitFailure,
	UnreachableError:   ExitFailure,
	UnknownOpcodeError: ExitFailure,
	NoEntryPoint:       ExitFailure,
	ArithmeticOverflow: ExitFailure,
}

func (c *cpu) SetTrapMessage(message string) {
	c.tm = message
}

func (c *cpu) Trap(n int) {
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

func (c *cpu) TrapAnonymous(n int) {
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
