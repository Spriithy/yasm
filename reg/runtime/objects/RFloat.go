package obj

import (
	"fmt"
	"unsafe"
)

// RFloat is the internal representation of a 64-bits floating point number
type RFloat float64

// ValueOf returns the value of the float
func (r *RFloat) ValueOf() interface{} {
	return float64(*r)
}

// TypeOf returns the float type constant
func (r *RFloat) TypeOf() RType {
	return TFloat
}

func (r *RFloat) String() RString {
	s := fmt.Sprintf("%d", r)
	return StringOf(s)
}

// SizeOf returns the memory size used by the float
func (r *RFloat) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
