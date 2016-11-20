package builtins

// An Object is the base type wrapper
type Object interface {
	// Class returns a string describing the object's class
	Class() string

	// String returns a printable representation of the object
	String() string

	// Equals test equality between two objects
	Equals(Object) bool

	// Copy returns an identical object yet somewhere else in memory
	Copy() Object
}
