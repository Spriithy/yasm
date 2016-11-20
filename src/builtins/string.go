package builtins

import (
	"fmt"
	"strings"
)

// The String type is exactly what you think it is
type String string

// Contains returns wether the string given is contained in the
// parent string
func (s String) Contains(s2 String) bool {
	return strings.Contains(string(s), string(s2))
}

// StartsWith reports wether the parent string starts with the
// given string
func (s String) StartsWith(s2 String) bool {
	return strings.HasPrefix(string(s), string(s2))
}

// EndsWith reports wether the parent string ends with the
// given string
func (s String) EndsWith(s2 String) bool {
	return strings.HasSuffix(string(s), string(s2))
}

// IndexOf returns the first apparition index of the given character
// in the string, unless it isn't found and returns -1 instead.
func (s String) IndexOf(c Char) Int {
	for i, c2 := range s {
		if Char(c2) == c {
			return Int(i)
		}
	}
	return Int(-1)
}

// LastIndexOf returns the last apparition index of the given character
// in the string, unless it isn't found and returns -1 instead.
func (s String) LastIndexOf(c Char) Int {
	l := -1
	for i, c2 := range s {
		if Char(c2) == c {
			l = i
		}
	}
	return Int(l)
}

// Size simply returns th length of the string
func (s String) Size() Int {
	return Int(len(s))
}

// Bytes returns the Char Array representing the string
func (s String) Bytes() []Char {
	return []Char(s)
}

// Format returns a string version of the String object but replacing
// escape sequences with their two character representation and other
// modifications ... to make the string more readable for humans in
// debug or inspection.
func (s String) Format() string {
	return fmt.Sprintf("%#v", s)
}

// TODO : Implement
//     -> Substring
//     -> Trim
//     -> Split

// --- Object Interface ---

// Class returns a string describing the object's class
func (String) Class() string {
	return "string"
}

// String returns a printable representation of the object
func (s String) String() string {
	return string(s)
}

// Equals test equality between two objects
// In the Null object case, nothing is equal to null
// but itself
func (s String) Equals(o Object) bool {
	return o == s
}

// Copy returns an identical object yet somewhere else in memory
func (s String) Copy() Object {
	return String(s)
}
