package yasm

import (
	"fmt"
	"os"
)

type Module struct {
	Name         string
	Libs         map[string]*Module
	Funcs        map[string]uintptr
	Instructions []Opcode
}

func (m *Module) RegisterFunc(name string, offset uintptr) {
	if addr, exists := m.Funcs[name]; exists {
		fmt.Printf("%s: redefinition of function '%s'. Previous declaration was here 0x%x\n", m.Name, name, addr)
		os.Exit(ExitFailure)
	}
	m.Funcs[name] = offset
}
