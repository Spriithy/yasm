package runtime

type vm struct {
	pc  int
	esp *reg
	tp  *reg
	me  *reg
	eax *reg
	ebx *reg
	ecx *reg
	edx *reg
}

func VirtualMachine() *vm {
	v := new(vm)
	v.pc = 0
	v.esp = makeReg()
	v.tp = makeReg()
	v.me = makeReg()
	v.eax = makeReg()
	v.ebx = makeReg()
	v.ecx = makeReg()
	v.edx = makeReg()
	return v
}
