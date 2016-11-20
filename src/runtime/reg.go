package runtime

import (
	"github.com/Spriithy/Polaroid/src/builtins"
)

type reg struct {
	t string
	v builtins.Object
}

// makeReg creates a new null-register
func makeReg() *reg {
	return &reg{"null", builtins.NULL}
}

func (r *reg) set(to builtins.Object) {
	r.t = to.Class()
	r.v = to
}

func (r *reg) get() builtins.Object {
	return r.v
}

func (r *reg) ptr() int {
	return int(r.int())
}

func (r *reg) char() builtins.Char {
	return r.v.(builtins.Char)
}

func (r *reg) int() builtins.Int {
	return r.v.(builtins.Int)
}

func (r *reg) flt() builtins.Float {
	return r.v.(builtins.Float)
}

func (r *reg) bool() bool {
	return r.v.Equals(builtins.TRUE)
}
