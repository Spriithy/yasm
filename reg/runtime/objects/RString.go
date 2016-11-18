package obj

import (
	"sync"
)

// RString is the internal representation of a string
type RString struct {
	size RInt

	sync.RWMutex
	bytes []byte
}

// StringOf returns a new instance of a string object
func StringOf(src string) RString {
	return RString{size: RInt(len(src)), bytes: []byte(src)}
}

// Get : Returns the byte of the stirng at the given index
// Returns an ErrOutOfBounds if index is out of bounds
func (r *RString) Get(idx RInt) (RByte, error) {
	r.RLock()
	defer r.RUnlock()
	if idx < 0 || idx >= r.size {
		return 0, ErrOutOfBounds
	}
	return RByte(r.bytes[idx]), nil
}

// Set : Modifies the byte at the given index
// Returns an ErrOutOfBounds if index is out of bounds
func (r *RString) Set(idx RInt, b RByte) error {
	r.Lock()
	defer r.Unlock()
	if idx < 0 || idx >= r.size {
		return ErrOutOfBounds
	}
	r.bytes[idx] = byte(b)
	return nil
}

// ValueOf returns a string representation of the string
func (r *RString) ValueOf() interface{} {
	return string(r.bytes)
}

// TypeOf returns the string type constant
func (r *RString) TypeOf() RType {
	return TString
}

func (r *RString) String() RString {
	r.Lock()
	defer r.Unlock()
	return StringOf(string(r.bytes))
}

// SizeOf returns the memory size used by the string instance
func (r *RString) SizeOf() RInt {
	return r.size
}
