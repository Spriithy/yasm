package yasm

const (
	rax = iota
	rbx
	rcx
	rdx
	rex
	rr0
	rr1
	rr2
	rr3
	rr4
	rr5
	rk0
	rk1
	rsp
	rfp
	rpc
)

var regName = map[int]string{
	0:  "%ax",
	1:  "%bx",
	2:  "%cx",
	3:  "%dx",
	4:  "%ex",
	5:  "%r0",
	6:  "%r1",
	7:  "%r2",
	8:  "%r3",
	9:  "%r4",
	10: "%r5",
	11: "%k0",
	12: "%k1",
	13: "%sp",
	14: "%fp",
	15: "%pc",
}
