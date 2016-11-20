package builtins

import "strconv"

// The Int type is the classic integer representation of numbers
type Int int64

// Class returns a string describing the object's class
func (Int) Class() string {
	return "int"
}

// String returns a printable representation of the object
func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Equals test equality between two objects
func (i Int) Equals(o Object) bool {
	return i == o
}

// Copy returns an identical object yet somewhere else in memory
// Numbers also don't follow the exact rule here
func (i Int) Copy() Object {
	return i
}
