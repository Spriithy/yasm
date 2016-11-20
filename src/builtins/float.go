package builtins

import (
	"strconv"
)

// The Float type is used to represented 64-bits floating point number
type Float float64

// Class returns a string describing the object's class
func (Float) Class() string {
	return "float"
}

// String returns a printable representation of the object
func (f Float) String() string {
	return strconv.FormatFloat(float64(f), 'f', 6, 64)
}

// Equals test equality between two objects
func (f Float) Equals(o Object) bool {
	return f == o
}

// Copy returns an identical object yet somewhere else in memory
// Numbers also don't follow the exact rule here
func (f Float) Copy() Object {
	return f
}
