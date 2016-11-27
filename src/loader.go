package small

import (
	"io/ioutil"
	"os"

	"github.com/Spriithy/go-colors"
)

// WordSize is the actual size of a word in a binary file
// of the SMALL language.
const WordSize = 1 << 5

// The Loader structure is used to load the bytecode
// binary files so the interpreter can execute it.
// The Loader provides header-checking for compiled
// files.
type loader struct {
	bytes []byte
	code  []Instruction
}

// Load is used to efficiently load a binary filed
// assembled for the SMALL vm
func Load(path string) []Instruction {
	l := new(loader)
	f, err := os.Open(path)
	if err != nil {
		println("small:", "couldn't open \""+path+"\".\n", colors.Red(colors.None, err))
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		println("small:", "couldn't read \""+path+"\".\n", colors.Red(colors.None, err))
		os.Exit(1)
	}

	l.feed(bytes)
	l.process()

	return l.code
}

func (l *loader) feed(bytes []byte) {
	l.bytes = bytes
}

func (l *loader) process() {
	l.code = make([]Instruction, len(l.bytes)/4)
	for at := 0; at < len(l.bytes)/4; at += 4 {
		b0, b1, b2, b3 := l.bytes[at], l.bytes[at+1], l.bytes[at+2], l.bytes[at+3]
		l.code[at/4] = ReadInstruction(b0, b1, b2, b3)
	}
}
