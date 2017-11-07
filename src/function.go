package yasm

type Function struct {
	Name   string
	Module *Module
	Caller *Function
	pc     uintptr
}
