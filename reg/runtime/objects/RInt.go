package obj

import (
	"fmt"
	"unsafe"
)

// RInt is the internal representation of a signed integer based on system's int
type RInt int

// ValueOf returns the value of the int
func (r *RInt) ValueOf() interface{} {
	return int(*r)
}

// TypeOf returns the int type constant
func (r *RInt) TypeOf() RType {
	return TInt
}

func (r *RInt) String() RString {
	s := fmt.Sprintf("%d", r)
	return StringOf(s)
}

// SizeOf returns the memory size used by the int
func (r *RInt) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
