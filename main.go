package main

import (
	"os"

	"github.com/Spriithy/SMALL/src"
)

func add(f *os.File, i uint32) {
	f.Write([]byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
}

func main() {
	i := []small.Instruction{
		small.CompileInstruction(0x0, 0x0, small.MOV, 11, 7, 0x0),
		small.CompileInstruction(0x1, 0x0, small.ADD, 3, 1, -0xf),
		small.CompileInstruction(0x1, 0x0, small.JMP, 3, 1, -0x0a0b),
	}

	small.Disasemble(i)
}
