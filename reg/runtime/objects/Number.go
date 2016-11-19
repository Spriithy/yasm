package objects

import (
	"errors"
	"fmt"
	"math"
)

// Numeric is used to denote any numeric type
type Numeric byte

const (
	// TypeByte is the byte type numeric constant
	TypeByte = Numeric(iota)

	// TypeInt is the int type numeric constant
	TypeInt

	// TypeUInt is the unsigned int type numeric constant
	TypeUInt

	// TypeFloat is the float type numeric constant
	TypeFloat
)

// Number is the internal representation for a Number
type Number struct {
	T Numeric // The Numeric type being stored
	b byte
	i int64
	u uint64
	f float64
}

// Byte wraps a byte into a runtime Number
func Byte(b byte) *Number {
	return &Number{T: TypeByte, b: b}
}

// Int wraps an integer into a runtime Number
func Int(i int64) *Number {
	return &Number{T: TypeInt, i: i}
}

// UInt wraps an unsigned integer into a runtime Number
func UInt(u uint64) *Number {
	return &Number{T: TypeUInt, u: u}
}

// Float wraps a floating point number into a runtime Number
func Float(f float64) *Number {
	return &Number{T: TypeFloat, f: f}
}

// Class returns the numeric name type of the Number
func (n *Number) Class() string {
	switch n.T {
	case TypeByte:
		return "byte"
	case TypeInt:
		return "int"
	case TypeUInt:
		return "uint"
	case TypeFloat:
		return "float"
	default:
		return "number"
	}
}

// Byte is used to recover the byte value of the number uncast
func (n *Number) Byte() (byte, error) {
	if n.T != TypeByte {
		return 0, errors.New("number is not byte")
	}
	return n.b, nil
}

// Int is used to recover the int value of the number uncast
func (n *Number) Int() (int64, error) {
	if n.T != TypeInt {
		return 0, errors.New("number is not int")
	}
	return n.i, nil
}

// UInt is used to recover the unsigned int value of the number uncast
func (n *Number) UInt() (uint64, error) {
	if n.T != TypeUInt {
		return 0, errors.New("number is not uint")
	}
	return n.u, nil
}

// Float is used to recover the float value of the number uncast
func (n *Number) Float() (float64, error) {
	if n.T != TypeByte {
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
func (n *Number) Cast(t Numeric) {
	switch t {
	case TypeByte:
		switch n.T {
		case TypeInt:
			n.b = byte(n.i % math.MaxUint8)
			n.i = 0
		case TypeUInt:
			n.b = byte(n.u % math.MaxUint8)
			n.u = 0
		case TypeFloat:
			n.b = byte(n.f)
			n.f = 0
		default:
		}
	case TypeInt:
		switch n.T {
		case TypeByte:
			n.i = int64(n.b)
			n.b = 0
		case TypeUInt:
			n.i = int64(n.u)
			n.u = 0
		case TypeFloat:
			n.i = int64(n.f)
			n.f = 0
		default:
		}
	case TypeUInt:
		switch n.T {
		case TypeByte:
			n.u = uint64(n.b)
			n.b = 0
		case TypeInt:
			n.u = uint64(n.i)
			n.i = 0
		case TypeFloat:
			n.u = uint64(n.f)
			n.f = 0
		default:
		}
	case TypeFloat:
		switch n.T {
		case TypeByte:
			n.f = float64(n.b)
			n.b = 0
		case TypeInt:
			n.f = float64(n.i)
			n.i = 0
		case TypeUInt:
			n.f = float64(n.u)
			n.u = 0
		default:
		}
	}
	n.T = t
}

// CastCopy allows the virtual machine to use the value as another type
// without losing its original information
func (n *Number) CastCopy(t Numeric) *Number {
	nn := n.Copy().(*Number)
	nn.Cast(t)
	return nn
}

func (n *Number) String() string {
	switch n.T {
	case TypeByte:
		return fmt.Sprintf("%d", n.b)
	case TypeInt:
		return fmt.Sprintf("%d", n.i)
	case TypeUInt:
		return fmt.Sprintf("%d", n.u)
	case TypeFloat:
		return fmt.Sprintf("%f", n.f)
	default:
		return "0"
	}
}
