package obj

import (
	"fmt"
	"unsafe"
)

// RByte is the internal representation of a byte
type RByte byte

// ValueOf retuns the value of r
func (r *RByte) ValueOf() interface{} {
	return byte(*r)
}

// TypeOf returns the byte type constant
func (r *RByte) TypeOf() RType {
	return TByte
}

func (r *RByte) String() RString {
	s := fmt.Sprintf("%d", r)
	return StringOf(s)
}

// SizeOf returns the memory size used by the byte
func (r *RByte) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
