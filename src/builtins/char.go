package builtins

import "strconv"

// Char is the natural type used to represent characters
// It is perfectly equivalent to an unsigned 8 bit integer
type Char uint8

// Class returns a string describing the object's class
func (Char) Class() string {
	return "char"
}

// String returns a printable representation of the object
func (c Char) String() string {
	return strconv.FormatInt(int64(c), 10)
}

// Equals test equality between two objects
func (c Char) Equals(o Object) bool {
	return c == o
}

// Copy returns an identical object yet somewhere else in memory
// Numbers also don't follow the exact rule here
func (c Char) Copy() Object {
	return c
}
