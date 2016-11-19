package objects

import (
	"errors"
	"sync"
)

// Array is the natural array data-structure with a finite amount
// of elements.
type Array struct {
	size int
	of   string
	sync.RWMutex
	array []Any
}

// CreateArray simply creates a new instance of an Array
// internal object with its Read/Write Mutex lock.
func CreateArray(size int, of string) (*Array, error) {
	if size < 1 {
		return nil, errors.New("array size must be a positive integer")
	}

	a := &Array{size: size, of: of, array: make([]Any, size)}
	e, err := NewBuiltin(of)
	if err != nil {
		return nil, err
	}

	for i := 0; i < size; i++ {
		a.array[i] = e
	}

	return a, nil
}

// Class is used to return the class name of any object
// Array Class name depends on the data being stored in it
func (a *Array) Class() string {
	if a.size != 0 {
		return a.array[0].Class() + "[]"
	}
	return "array[]"
}

// Copy returns a Copy of the current Array
func (a *Array) Copy() Any {
	return &Array{size: a.size, array: a.array}
}

// Read is used to retrieve a value from the Array
// It returns an "index out of bounds" error if the index is
// either negative or larger than the Array's size
func (a *Array) Read(idx int) (*Any, error) {
	if idx < 0 || idx >= a.size {
		return nil, errors.New("array index out of bounds")
	}

	a.RLock()
	defer a.RUnlock()
	return &a.array[idx], nil
}

// Write is used to update a value being stored in one
// of the Array's cells. Types must match
func (a *Array) Write(idx int, val Any) error {
	if idx < 0 || idx >= a.size {
		return errors.New("array index out of bounds")
	}

	a.Lock()
	defer a.Unlock()
	valc := val.Class()

	if valc != a.of && a.of != "null" {
		return errors.New("array type mismatch expected: " + a.of + " got " + valc)
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
