package objects

import (
	"errors"
	"strings"
	"sync"
)

// Array is the natural array data-structure with a finite amount
// of elements.
type Array struct {
	size int
	sync.RWMutex
	array []Any
}

// CreateArray simply creates a new instance of an Array
// internal object with its Read/Write Mutex lock.
func CreateArray(size int, def Any) (*Array, error) {
	if size < 1 {
		return nil, errors.New("array size must be a positive integer")
	}

	a := &Array{size: size, array: make([]Any, size)}
	for i := 0; i < size; i++ {
		a.array[i] = def.Copy()
	}

	return a, nil
}

func (a *Array) Class() string {
	if a.size != 0 {
		return "array[" + a.array[0].Class() + "]"
	}
	return "array[]"
}

func (a *Array) Copy() Any {
	return &Array{size: a.size, array: a.array}
}

func (a *Array) Zero() Any {
	return &Array{size: 1, array: []Any{Null{}}}
}

func (a *Array) Read(idx int) (*Any, error) {
	if idx < 0 || idx >= a.size {
		return nil, errors.New("array index out of bounds")
	}

	a.RLock()
	defer a.RUnlock()
	return &a.array[idx], nil
}

func (a *Array) Write(idx int, val Any) error {
	if idx < 0 || idx >= a.size {
		return errors.New("array index out of bounds")
	}

	a.Lock()
	defer a.Unlock()
	valc := val.Class()
	exp := a.array[idx].Class()

	if valc != exp && !(strings.HasPrefix(valc, "array") && strings.HasPrefix(valc, "array")) {
		return errors.New("array type mismatch expected: " + exp + " got " + valc)
	}

	a.array[idx] = val.Copy()
	return nil
}

func (a *Array) String() string {
	if len(a.array) == 0 {
		return "[]"
	}

	str := ""
	a.RLock()
	for _, pair := range a.array {
		str += ", " + pair.String()
	}
	a.RUnlock()
	return "[" + str[2:] + "]"
}
