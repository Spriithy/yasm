package builtins

import "github.com/Spriithy/Polaroid/old/utils"

type Byte rune

func (b Byte) Class() string {
	return "{Byte hash=" + b.Hash().String() + "}"
}

func (b Byte) TypeOf() BuiltinType {
	return BUILTIN_BYTE
}

func (b Byte) String() string {
	s, l := format("%X", rune(b))
	return "0x" + s[1:l - 1]
}

func (b Byte) Hash() utils.Hash {
	return utils.Hash(b)
}

func (b Byte) Copy() Object {
	return Byte(b)
}