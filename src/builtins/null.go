package builtins

// NULL must be the only NULL object used
var NULL = null{}

// Null is the intuitive object desribing the absence of data
type null struct{}

// Class returns a string describing the object's class
func (null) Class() string {
	return "null"
}

// String returns a printable representation of the object
func (null) String() string {
	return "null"
}

// Equals test equality between two objects
// In the Null object case, nothing is equal to null
// but itself
func (null) Equals(o Object) bool {
	return o == NULL
}

// Copy returns an identical object yet somewhere else in memory
// Only one Null value is to be used
func (n null) Copy() Object {
	return n
}
