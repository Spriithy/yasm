package builtins

import "github.com/Spriithy/Polaroid/old/utils"

type Null int

const NEW_NULL Null = Null(0)

func (n Null) Class() string {
	return "{Null}"
}

func (n Null) TypeOf() BuiltinType {
	return BUILTIN_NULL
}

func (n Null) String() string {
	return "null"
}

func (n Null) Hash() utils.Hash {
	return utils.Hash(0)
}

func (n Null) Copy() Object {
	return Null(0)
}

