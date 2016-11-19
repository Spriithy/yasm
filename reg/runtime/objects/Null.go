package objects

// Null is used to represent the absence of value
type Null struct{}

// Class returns the class name of any object
func (n Null) Class() string {
	return "null"
}

// String returns the "null" string
func (n Null) String() string {
	return "null"
}

// Copy returns a copy of the Null Object
func (n Null) Copy() Any {
	return Null{}
}
