package builtins

import (
	"sync"

	"github.com/Spriithy/Polaroid/src/errors"
	"github.com/Spriithy/go-uuid"
)

// Array is the natural array structure with a single element type supported
type Array struct {
	id uuid.UUID
	of string

	sync.RWMutex
	els []Object
}

// CreateArray will instantiate a new Array object with its associated type and initial length
// Every element will be instantiated as NULL
func CreateArray(size int, class string) *Array {
	if size < 0 {
		errors.Report("array size must be a positive integer")
	}

	if len(class) < 1 {
		errors.Report("invalid array element type: " + class)
	}

	a := new(Array)
	a.id = uuid.NextUUID()
	a.of = class
	a.els = make([]Object, size)

	a.Lock()
	for i := range a.els {
		a.els[i] = NULL
	}
	a.Unlock()

	return a
}

// Set is used to update an Array's elements. If the index is out of the array's bounds
// the execution will halt and report an error. Same if type mismatch.
func (a *Array) Set(at int, o Object) {
	if o.Class() != a.of {
		errors.Report("array element type mismatch. Expected " + a.of + ", got " + o.Class())
	}

	switch {
	case at < len(a.els) && at >= 0:
		a.Lock()
		a.els[at] = o
		a.Unlock()
	default:
		errors.Report("index array out of bounds")
	}
}

// Get is used to retrieve an Array's elements
func (a *Array) Get(at int) Object {
	switch {
	case at < len(a.els) && at >= 0:
		a.RLock()
		defer a.RUnlock()
		return a.els[at]
	default:
		errors.Report("index array out of bounds")
		return nil
	}
}

// Size returns the length of the array
func (a *Array) Size() Int {
	return Int(len(a.els))
}

// --- Object Interface ---

// Class returns a string describing the object's class
func (a *Array) Class() string {
	return a.of + "[]"
}

// String returns a printable representation of the object
func (a *Array) String() string {
	if len(a.els) == 0 {
		return "[]"
	}

	str := ""
	a.RLock()
	for _, el := range a.els {
		if el.Class() == "string" {
			str += ", " + el.(String).Format()
		} else {
			str += ", " + el.String()
		}
	}
	a.RUnlock()

	return "[" + str[2:] + "]"
}

// Equals test equality between two objects
func (a *Array) Equals(o Object) bool {
	if a.Class() != o.Class() {
		return false
	}
	return a.id == o.(*Array).id
}

// Copy returns an identical object yet somewhere else in memory
func (a *Array) Copy() Object {
	aa := &Array{id: uuid.NextUUID(), of: a.of, els: make([]Object, len(a.els))}
	a.RLock()
	aa.Lock()
	for i, v := range a.els {
		aa.els[i] = v.Copy()
	}
	aa.Unlock()
	a.RUnlock()
	return aa
}
