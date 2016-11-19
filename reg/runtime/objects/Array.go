package objects

import (
	"errors"
	"sync"
)

type Array struct {
	size int
	sync.RWMutex
	array []Any
}

// CreateArray simply creates a new instance of an Array
// internal object with its Read/Write Mutex lock.
func CreateArray(size int) (*Array, error) {
	if size < 1 {
		return nil, errors.New("array size must be positive")
	}

	a := &Array{size: size, array: make([]Any, size)}
	for i := 0; i < size; i++ {
		a.array[i] = Null{}
	}

	return a, nil
}

func (a *Array) Class() string {
	return "array"
}

func (a *Array) Copy() Any {
	return &Array{size: a.size, array: a.array}
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
