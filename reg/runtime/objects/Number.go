package objects

import (
	"fmt"
)

// An Int is the internal representation of an integer
type Int int64

// Class returns the integer's type
func (i Int) Class() string {
	return "int"
}

// Copy returns the very same integer
func (i Int) Copy() Any {
	return Int(i)
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

// A Float is the internal representation of a floating point number
type Float float64

// Class returns the float's type
func (f Float) Class() string {
	return "float"
}

// Copy returns the very same float number
func (f Float) Copy() Any {
	return Float(f)
}

func (f Float) String() string {
	return fmt.Sprintf("%f", f)
}
