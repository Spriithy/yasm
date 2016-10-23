package builtins

import (
	"github.com/Spriithy/Polaroid/src/utils"
	"fmt"
)

type Object interface {
	Class() string
	TypeOf() BuiltinType
	String() string
	Hash() utils.Hash
	Copy() Object
}

func format(f string, a ... interface{}) (string, int) {
	s := fmt.Sprintf(f, a)
	return s, len(s)
}
