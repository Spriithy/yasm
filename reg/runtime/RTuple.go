package runtime

import (
	"sync"
	"unsafe"
)

// RTuple is the internal representation of a Tuple
// A Tuple is an ordered set of objects stored with integer keys
type RTuple struct {
	size RInt

	sync.RWMutex
	els []RObject
}

// Get : Returns the Object stored at the given index in the Tuple
// Returns an ErrOutOfBounds if index is out of bounds
func (r *RTuple) Get(idx RInt) (RObject, error) {
	r.RLock()
	defer r.RUnlock()
	if idx < 0 || idx >= r.size {
		return nil, ErrOutOfBounds
	}
	return r.els[idx], nil
}

// Set : Modifies the tuple index-th element
// Returns ErrOutOfBounds if the index is out of bounds
func (r *RTuple) Set(idx RInt, o RObject) error {
	r.Lock()
	defer r.Unlock()
	if idx < 0 || idx >= r.size {
		return ErrOutOfBounds
	}
	r.els[idx] = o
	return nil
}

// Append is used to append an Object to the tuple
func (r *RTuple) Append(o RObject) {
	r.Lock()
	defer r.Unlock()
	r.size++
	r.els = append(r.els, o)
}

// Remove is used to remove an entry in the tuple, and resize it accordingly
// Returns an ErrOutOfBounds is index is out of bounds
func (r *RTuple) Remove(idx RInt) error {
	r.RLock()
	defer r.RUnlock()
	if idx < 0 || idx >= r.size {
		return ErrOutOfBounds
	}
	r.els = append(r.els[:idx], r.els[idx+1:]...)
	return nil
}

// ValueOf doesn't return anything intersting for now
func (r *RTuple) ValueOf() interface{} {
	return 0
}

// TypeOf returns the tuple type constant
func (r *RTuple) TypeOf() RType {
	return TTuple
}

func (r *RTuple) String() RString {
	str := ""
	r.Lock()
	defer r.Unlock()
	for _, el := range r.els {
		str += ", " + string(el.String().bytes)
	}
	str = str[2:]
	return StringOf("(" + str + ")")
}

// SizeOf returns the memory size used by the tuple
func (r *RTuple) SizeOf() RInt {
	return RInt(unsafe.Sizeof(r))
}
