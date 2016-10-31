package builtins

import "github.com/Spriithy/Polaroid/old/utils"

type Array struct {
	data []Object
	len  Integer
}

func NewArray(fcount int) *Array {
	a := new(Array)
	a.data = make([]Object, fcount)
	a.len = Integer(fcount)
	for i := 0; i < fcount; i++ {
		a.data[i] = NEW_NULL
	}
	return a
}

////////////////////////////////////////////////////////////////////////////////
//
// Array Builtin functions
//
////////////////////////////////////////////////////////////////////////////////

func (a Array) Class() string {
	t := string(a.TypeOf())
	h := a.Hash().String()
	return "{" + t + " hash=" + h + "}"
}

func (a Array) TypeOf() BuiltinType {
	return a.data[0].TypeOf() + "[]"
}

func (a Array) String() string {
	str := ""
	for _, e := range a.data {
		if e.TypeOf().IsNative() {
			str += e.String() + ", "
		} else {
			str += e.Class() + ", "
		}
	}
	return string(a.TypeOf()) + "{" + str[:len(str) - 2] + "}"
}

func (a Array) Hash() utils.Hash {
	var h utils.Hash = 0
	for i := 0; i < len(a.data); i++ {
		h += a.data[i].Hash()
	}
	return h + utils.HashString(string(a.TypeOf()))
}

func (a Array) Copy() Object {
	data := make([]Object, a.len)
	copy(data, a.data)
	return Array{data, a.len}
}

func (a Array) Length() Integer {
	return a.len
}

func (a Array) Read(at Integer) (Object, error) {
	if at < 0 || at > Integer(len(a.data)) {
		return nil, &ArrayIndexOutOfBoundsError{
			msg:    "Trying to read array index out of bounds (%d, %d), array: *%s",
			at:     at,
			of:     Integer(len(a.data)),
			a:      a,
		}
	}
	return a.data[at], nil
}

func (a Array) Write(at Integer, o Object) error {
	if at < 0 || at > Integer(len(a.data)) {
		return &ArrayIndexOutOfBoundsError{
			msg:    "Trying to index array out of bounds (%d, %d), array: *%s",
			at:     at,
			of:     Integer(len(a.data)),
			a:      a,
		}
	}
	dst := a.data[at]
	if dst.TypeOf() != o.TypeOf() {
		return &TypeMismatchError{
			msg:    "Cannot write type %s where type %s is expected in array: *%s",
			t0:     o.TypeOf(),
			t1:     dst.TypeOf(),
			o:      a,
		}
	}
	a.data[at] = o
	return nil
}

func (a Array) WriteNew(at Integer, o Object) error {
	if at < 0 || at > Integer(len(a.data)) {
		return &ArrayIndexOutOfBoundsError{
			msg:    "Trying to index array out of bounds (%d, %d), array: *%s",
			at:     at,
			of:     Integer(len(a.data)),
			a:      a,
		}
	}
	a.data[at] = o
	return nil
}