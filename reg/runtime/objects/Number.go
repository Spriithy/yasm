package objects

import (
	"errors"
	"fmt"
)

// NumericType is used to denote any numeric type
type NumericType byte

const (
	// ByteType is the byte type numeric constant
	ByteType = NumericType(iota)

	// IntType is the int type numeric constant
	IntType

	// UIntType is the unsigned int type numeric constant
	UIntType

	// FloatType is the float type numeric constant
	FloatType
)

// Number is the internal representation for a Number
type Number struct {
	T NumericType // The Numeric type being stored
	b byte
	i int64
	u uint64
	f float64
}

// Byte wraps a byte into a runtime Number
func Byte(b byte) *Number {
	return &Number{T: ByteType, b: b}
}

// Int wraps an integer into a runtime Number
func Int(i int64) *Number {
	return &Number{T: IntType, i: i}
}

// UInt wraps an unsigned integer into a runtime Number
func UInt(u uint64) *Number {
	return &Number{T: UIntType, u: u}
}

// Float wraps a floating point number into a runtime Number
func Float(f float64) *Number {
	return &Number{T: FloatType, f: f}
}

// Class returns the numeric name type of the Number
func (n *Number) Class() string {
	switch n.T {
	case ByteType:
		return "byte"
	case IntType:
		return "int"
	case UIntType:
		return "uint"
	case FloatType:
		return "float"
	default:
		return "number"
	}
}

// Zero returns the 0-value of the number's type
func (n *Number) Zero() Any {
	switch n.T {
	case ByteType:
		return Byte(0)
	case IntType:
		return Int(0)
	case UIntType:
		return UInt(0)
	case FloatType:
		return Float(0)
	}
	return Int(0)
}

// Byte is used to recover the byte value of the number uncast
func (n *Number) Byte() (byte, error) {
	if n.T != ByteType {
		return 0, errors.New("number is not byte")
	}
	return n.b, nil
}

// Int is used to recover the int value of the number uncast
func (n *Number) Int() (int64, error) {
	if n.T != IntType {
		return 0, errors.New("number is not int")
	}
	return n.i, nil
}

// UInt is used to recover the unsigned int value of the number uncast
func (n *Number) UInt() (uint64, error) {
	if n.T != UIntType {
		return 0, errors.New("number is not uint")
	}
	return n.u, nil
}

// Float is used to recover the float value of the number uncast
func (n *Number) Float() (float64, error) {
	if n.T != ByteType {
		return 0, errors.New("number is not float")
	}
	return n.f, nil
}

// Copy returns a naive copy of the source Number
func (n *Number) Copy() Any {
	return &Number{n.T, n.b, n.i, n.u, n.f}
}

// Cast allows the virtual machine to cast the Number into another
// Numeric type.
func (n *Number) Cast(t NumericType) {
	switch t {
	case ByteType:
		n.f = float64(n.b)
		n.b = 0
	case IntType:
		n.f = float64(n.i)
		n.i = 0
	case UIntType:
		n.f = float64(n.u)
		n.u = 0
	default:
	}
	n.T = t
}

// CastCopy allows the virtual machine to use the value as another type
// without losing its original information
func (n *Number) CastCopy(t NumericType) *Number {
	nn := n.Copy().(*Number)
	nn.Cast(t)
	return nn
}

// Returns the print-friendly version of the number being stored
func (n *Number) String() string {
	switch n.T {
	case ByteType:
		return fmt.Sprintf("%d", n.b)
	case IntType:
		return fmt.Sprintf("%d", n.i)
	case UIntType:
		return fmt.Sprintf("%d", n.u)
	case FloatType:
		return fmt.Sprintf("%f", n.f)
	default:
		return "0"
	}
}
