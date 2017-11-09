package yasm

type Function struct {
	Name   string
	Module *Module
	Caller *Function
	pc     uintptr
}

func (c *cpu) makeFunc(caller *Function, pc uintptr) {

}
