package objects

// Bool is the internal representation of a boolean value
type Bool bool

// Class is used to return the class name of any object
func (b Bool) Class() string {
	return "bool"
}

// Returns either true or false in a string
func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

// Copy returns a copy of the boolean
func (b Bool) Copy() Any {
	return Bool(b)
}
