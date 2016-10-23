package builtins

import (
	"github.com/Spriithy/Polaroid/src/utils"
	"math"
)

type Real float64

func (r Real) Class() string {
	return "{Real hash=" + r.Hash().String() + "}"
}

func (r Real) TypeOf() BuiltinType {
	return BUILTIN_REAL
}

func (r Real) String() string {
	s, l := format("%f", float64(r))
	return s[1:l - 1]
}

func (r Real) Hash() utils.Hash {
	v := math.Float64bits(float64(r))
	return utils.Hash(v ^ (v >> 32))
}

func (r Real) Copy() Object {
	return Real(r)
}
