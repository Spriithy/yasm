package builtins

// TRUE must be the only True constant in use
var TRUE = boolean(true)

// FALSE must be the only False constant in use
var FALSE = boolean(false)

// Boolean is the straight forward boolean logic values implementation
// with the restriction to not allocate more than one of each False and
// True values.
type boolean bool

// Class returns a string describing the object's class
func (boolean) Class() string {
	return "boolean"
}

// String returns a printable representation of the object
func (b boolean) String() string {
	if b {
		return "true"
	}
	return "false"
}

// Equals test equality between two objects
func (b boolean) Equals(o Object) bool {
	return b == o
}

// Copy returns an identical object yet somewhere else in memory
// Exception for booleans as only one instance of each must exist
func (b boolean) Copy() Object {
	return b
}
