package objects

import "errors"

func makeZeroArray() Any {
	a := &Array{size: 0, of: "null", array: make([]Any, 0)}
	return a
}

func makeZeroTable() Any {
	t := CreateTable()
	return t
}

var builtinZeros = map[string]Any{
	"null":   Null{},
	"bool":   Bool(false),
	"byte":   Byte(0),
	"int":    Int(0),
	"uint":   UInt(0),
	"float":  Float(0),
	"string": String(""),
	"array":  makeZeroArray(),
	"table":  makeZeroTable(),
}

// NewBuiltin is used to return a new class object of the given internal type
// that is builtin
func NewBuiltin(t string) (Any, error) {
	if z, ok := builtinZeros[t]; ok {
		return z, nil
	}
	return nil, errors.New("unknown builtin type: " + t)
}
