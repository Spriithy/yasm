package builtins

import "github.com/Spriithy/Polaroid/src/utils"

type Integer int64

func (i Integer) Class() string {
	return "{Integer hash=" + i.Hash().String() + "}"
}

func (i Integer) TypeOf() BuiltinType {
	return BUILTIN_INTEGER
}

func (i Integer) String() string {
	s, l := format("%d", i)
	return s[1:l - 1]
}

func (i Integer) Hash() utils.Hash {
	return utils.Hash(i)
}

func (i Integer) Copy() Object {
	return Integer(i)
}
