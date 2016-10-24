package runtime

import (
	"github.com/Spriithy/Polaroid/src/builtins"
	"os"
)

// Used to represent the Inner VM stack
type Stack struct {
	data []builtins.Object // Stack Data
	sp   int               // Stack Pointer
}

// Simply constructs a new Emty Stack
//     s.data     will be an empty slice of VirtValues
//     s.sp       is set to -1
func NewStack() *Stack {
	return &Stack{make([]builtins.Object, 1000), -1}
}

// Peeks the top of the stack without poping it
func (s *Stack) Peek() builtins.Object {
	return s.data[s.sp] // Simple look-up
}

func (s *Stack) Push(v... builtins.Object) {
	for _, k := range v {
		s.sp++
		if len(s.data) == s.sp {
			(*s).data = append((*s).data, k)
			continue
		}
		s.data[s.sp] = k
	}
}

func (s *Stack) PopBoolean() builtins.Boolean {
	v := s.Pop()
	switch v.(type) {
	case builtins.Boolean:
		return v.(builtins.Boolean)
	case builtins.Null:
		println("NullPointerException: expected Boolean got Null")
		os.Exit(1)
	default:
		// TODO VM ERROR
		panic(v.Class())
	}
	return false
}

func (s *Stack) PopByte() builtins.Byte {
	v := s.Pop()
	switch v.(type) {
	case builtins.Byte:
		return v.(builtins.Byte)
	case builtins.Null:
		println("NullPointerException: expected Byte got Null")
		os.Exit(1)
	default:
		// TODO VM ERROR
		panic(v.Class())
	}
	return -1
}

func (s *Stack) PopInteger() builtins.Integer {
	v := s.Pop()
	switch v.(type) {
	case builtins.Integer:
		return v.(builtins.Integer)
	case builtins.Null:
		println("NullPointerException: expected Integer got Null")
		os.Exit(1)
	default:
		// TODO VM ERROR
		panic(v.Class())
	}
	return -1
}

func (s *Stack) PopReal() builtins.Real {
	v := s.Pop()
	switch v.(type) {
	case builtins.Real:
		return v.(builtins.Real)
	case builtins.Null:
		println("NullPointerException: expected Real got Null")
		os.Exit(1)
	default:
		// TODO VM ERROR
		panic(v.Class())
	}
	return -1
}

func (s *Stack) PopString() builtins.String {
	v := s.Pop()
	switch v.(type) {
	case builtins.String:
		return v.(builtins.String)
	case builtins.Null:
		println("NullPointerException: expected String got Null")
		os.Exit(1)
	default:
		// TODO VM ERROR
		panic(v.Class())
	}
	return builtins.String{}
}

func (s *Stack) Pop() builtins.Object {
	if s.Empty() {
		println("Cannot pop from empty stack!", s.String())
		os.Exit(1)
	}

	v := s.data[s.sp]
	s.sp--
	return v
}

func (s *Stack) Empty() bool {
	return s.sp < 0
}

func (s *Stack) String() string {
	if s.Empty() {
		return "[]"
	}
	str := ""
	for i := s.sp; i >= 0; i-- {
		str += " " + s.data[i].String()
	}
	return "[" + str[1:] + "]"
}