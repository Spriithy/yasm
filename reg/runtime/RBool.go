package runtime

import (
	"unsafe"
)

// RBool is the internal representation of a boolean value
type RBool bool

// ValueOf returns the value of r
func (r *RBool) ValueOf() interface{} {
	return bool(*r)
}

// TypeOf returns the boolean type constant
func (r *RBool) TypeOf() RType {
	return TBool
}

func (r *RBool) String() RString {
	if *r {
		return StringOf("true")
	}
	return StringOf("false")
}

// SizeOf returns the memory size used by the boolean
func (r *RBool) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
