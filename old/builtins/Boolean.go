package builtins

import "github.com/Spriithy/Polaroid/old/utils"

type Boolean bool

func (b Boolean) Class() string {
	return "{Boolean hash=" + b.Hash().String() + "}"
}

func (b Boolean) TypeOf() BuiltinType {
	return BUILTIN_BOOLEAN
}

func (b Boolean) String() string {
	if b {
		return "true"
	}
	return "false"
}

func (b Boolean) Hash() utils.Hash {
	if b {
		return utils.Hash(1)
	}
	return utils.Hash(0)
}

func (b Boolean) Copy() Object {
	return Boolean(b)
}
