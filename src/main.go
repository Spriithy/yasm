package main

import (
	"os"

	small "github.com/Spriithy/small/src"
)

func add(f *os.File, i uint32) {
	f.Write([]byte{byte(i), byte(i >> 8), byte(i >> 16), byte(i >> 24)})
}

func main() {
	i := small.CompileInstruction(0x0, 0x0, small.MOV, 11, 7, 3)
	small.Disasemble(i)
}
