package objects

import (
	"errors"
	"fmt"
	"sync"
)

// StringLit is the String Litteral object descriptor
type StringLit struct {
	size int
	sync.RWMutex
	bytes []byte
}

// String is used to create a new internal string
func String(src string) *StringLit {
	return &StringLit{size: len(src), bytes: []byte(src)}
}

// Class is used to return the class name of any object
func (s *StringLit) Class() string {
	return "string"
}

// Copy returns the internal copy of the string
func (s *StringLit) Copy() Any {
	return &StringLit{size: s.size, bytes: s.bytes}
}

// Read is used to read a character (byte) in the source string
// It returns an "Index out of bounds" error if the index is negative or
// larger than the string's size
func (s *StringLit) Read(at int) (byte, error) {
	if at >= s.size || at < 0 {
		return 0, errors.New("string index out of bounds")
	}

	s.RLock()
	defer s.RUnlock()

	return s.bytes[at], nil
}

// Write is used to update a character in the StringLit
// It returns an "Index out of bounds" error if the index is negative or
// larger than the string's size
func (s *StringLit) Write(at int, b byte) error {
	if at >= s.size || at < 0 {
		return errors.New("string index out of bounds")
	}

	s.Lock()
	defer s.Unlock()
	s.bytes[at] = b
	return nil
}

// Repr returns the string representation of the StringLit being stored
// It is the raw-displayable string version
func (s *StringLit) Repr() string {
	return string(s.bytes)
}

// String returns a human-readable with escapes codes etc. string version
// of the StringLit being stored
func (s *StringLit) String() string {
	s.RLock()
	defer s.RUnlock()
	return fmt.Sprintf("%#v", string(s.bytes))
}
