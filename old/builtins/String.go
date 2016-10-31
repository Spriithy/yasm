package builtins

import (
	"github.com/Spriithy/Polaroid/old/utils"
	"unicode/utf8"
	"strings"
)

type String struct {
	str string
}

func NewString(src string) *String {
	return &String{src}
}

////////////////////////////////////////////////////////////////////////////////
//
// String Builtin functions
//
////////////////////////////////////////////////////////////////////////////////

func (s String) Class() string {
	return "{String hash=" + s.Hash().String() + "}"
}

func (s String) TypeOf() BuiltinType {
	return BUILTIN_STRING
}

func (s String) String() string {
	return s.str
}

func (s String) Hash() utils.Hash {
	return utils.HashString(s.str)
}

func (s String) Copy() Object {
	return String{s.str}
}

// String Builtin functions

func (s String) Join(other Object) {
	s.str += other.String()
}

func (s String) Multiply(by Integer) {
	for ; by > 0; by-- {
		s.Join(s)
	}
}

func (s String) Reverse() {
	o := make([]rune, utf8.RuneCountInString(s.str));
	i := len(o);
	for _, c := range s.str {
		i--;
		o[i] = c;
	}
	s.str = string(o)
}

func (s String) Length() Integer {
	return Integer(len(s.str))
}

func (s String) Empty() Boolean {
	return len(s.str) == 0
}

func (s String) ByteAt(index Integer) Byte {
	return Byte(s.str[index])
}

func (s String) Bytes() []Byte {
	return []Byte(s.str)
}

func (s String) StartsWith(with String) Boolean {
	return Boolean(strings.HasPrefix(s.str, with.str))
}

func (s String) EndsWith(with String) Boolean {
	return Boolean(strings.HasSuffix(s.str, with.str))
}

func (s String) Contains(with String) Boolean {
	return Boolean(strings.Contains(s.str, with.str))
}

func (s String) Slice(from, to Integer) String {
	if from < 0 {
		return String{""}
	}

	if to < 0 {
		return s.Slice(-to + 1, from)
	}

	if from > to {
		return String{""}
	}

	return String{s.str[from:to]}
}