package builtins

import (
	"fmt"
)

type TypeMismatchError struct {
	msg    string
	t0, t1 BuiltinType
	o      Object
}

func (e TypeMismatchError) Error() string {
	return fmt.Sprintf(e.msg, e.t0, e.t1, e.o.Class())
}

type FieldIndexOutOfBoundsError struct {
	msg    string
	at, of Integer
	st     Structure
}

func (e FieldIndexOutOfBoundsError) Error() string {
	return fmt.Sprintf(e.msg, e.at, e.of, e.st.Class())
}

type ArrayIndexOutOfBoundsError struct {
	msg    string
	at, of Integer
	a      Array
}

func (e ArrayIndexOutOfBoundsError) Error() string {
	return fmt.Sprintf(e.msg, e.at, e.of, e.a.Class())
}
