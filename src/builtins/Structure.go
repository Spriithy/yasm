package builtins

import (
	"github.com/Spriithy/Polaroid/src/utils"
)

type Structure struct {
	name   String
	fields []Object
}

func NewStruct(name string, fcount int) *Structure {
	s := new(Structure)
	s.name = *NewString(name)
	s.fields = make([]Object, fcount)
	for i := 0; i < fcount; i++ {
		s.fields[i] = NEW_NULL
	}
	return s
}

////////////////////////////////////////////////////////////////////////////////
//
// Structure Builtin functions
//
////////////////////////////////////////////////////////////////////////////////

func (s Structure) Class() string {
	return "{" + s.name.str + " hash=" + s.Hash().String() + "}"
}

func (s Structure) TypeOf() BuiltinType {
	return BuiltinType(s.name.str)
}

func (s Structure) String() string {
	str := ""
	for _, f := range s.fields {
		if f.TypeOf().IsNative() {
			str += f.String() + ", "
		} else {
			str += f.Class() + ", "
		}
	}
	return s.name.str + "{" + str[:len(str) - 2] + "}"
}

func (s Structure) Hash() utils.Hash {
	var h utils.Hash = 0
	for i := 0; i < len(s.fields); i++ {
		h += s.fields[i].Hash()
	}
	return h + utils.HashString(s.name.str)
}

func (s Structure) Copy() Object {
	fields := make([]Object, len(s.fields))
	copy(fields, s.fields)
	return Structure{s.name, fields}
}

func (s Structure) Read(at Integer) Object {
	if at > Integer(len(s.fields)) || at < 0 {
		return NEW_NULL
	}
	return s.fields[at]
}

func (s Structure) Write(at Integer, o Object) error {
	if at < 0 || at > Integer(len(s.fields)) {
		return &FieldIndexOutOfBoundsError{
			msg: "Trying to index struct field out of bounds (%d, %d) in struct *%s",
			at: at,
			of: Integer(len(s.fields)),
			st: s,
		}
	}
	dst := s.fields[at]
	if dst.TypeOf() != o.TypeOf() {
		return &TypeMismatchError{
			msg:    "Cannot write type %s where type %s is expected in struct *%s",
			t0:     o.TypeOf(),
			t1:     dst.TypeOf(),
			o:      s,
		}
	}
	s.fields[at] = o
	return nil
}

func (s Structure) WriteNew(at Integer, o Object) error {
	if at < 0 || at > Integer(len(s.fields)) {
		return &FieldIndexOutOfBoundsError{
			msg: "Trying to index struct field out of bounds (%d, %d) in struct *%s",
			at: at,
			of: Integer(len(s.fields)),
			st: s,
		}
	}
	s.fields[at] = o
	return nil
}