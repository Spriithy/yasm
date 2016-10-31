package runtime

import (
	"fmt"
)

type List struct {
	len  int
	data []Object
}

func NewList(cap int) *List {
	assert(cap > 0, "initial list capacity must be > 0")
	l := new(List)
	l.len = cap
	l.data = make([]Object, cap)
	for i, _ := range l.data {
		l.data[i] = Wrap(nil)
	}
	return l
}

// Object Interface implementation

func (l List) ToString() string {
	return fmt.Sprintf("list[%d]", l.len)
}

func (l List) Equals(obj Object) bool {
	var val, ok = obj.(*List)
	if !ok {
		return false
	}

	for k, v := range l.data {
		if !v.Equals(val.data[k]) {
			return false
		}
	}

	return true
}

func (l List) Kind() ObjectKind {
	return ListKind
}

func (l List) Implements(k ObjectKind) bool {
	return k == l.Kind()
}

///////////

func (l *List) Length() int {
	return l.len
}

func (l List) Set(at int, el Object) bool {
	assert(at < l.len, "Trying to reach index out of list bounds!")
	l.data[at] = el
	return true
}

func (l List) Get(at int) Object {
	assert(at < l.len, "Trying to reach index out of list bounds!")
	return l.data[at]
}

func (l *List) Append(el ... Object) int {
	l.data = append(l.data, el ...)
	l.len += len(el)
	return l.len
}

func (l *List) SubList(from, to int) *List {
	assert(to > from && from >= 0, "Invalid sublist offsets (to > from or from < 0)")
	l2 := NewList(to - from)
	for k, _ := range l2.data {
		l2.data[k] = l.data[from + k]
	}
	return l2
}
