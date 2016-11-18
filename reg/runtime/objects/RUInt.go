package obj

import (
	"fmt"
	"unsafe"
)

// RUInt is the internal representation of an unsigned integer
type RUInt uint

// ValueOf returns the value of the uint
func (r *RUInt) ValueOf() interface{} {
	return uint(*r)
}

// TypeOf returns the uint type constant
func (r *RUInt) TypeOf() RType {
	return TUInt
}

func (r *RUInt) String() RString {
	s := fmt.Sprintf("%d", r)
	return StringOf(s)
}

// SizeOf returns the memory size used by the uint
func (r *RUInt) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
