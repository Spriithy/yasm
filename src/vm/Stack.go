package vm

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

func (s *Stack) Push(v builtins.Object) {
	s.sp++
	if len(s.data) == s.sp {
		(*s).data = append((*s).data, v)
		return
	}
	s.data[s.sp] = v
}

func (s *Stack) PushI(i int64) {
	s.Push(builtins.Integer(i))
}

func (s *Stack) PushB(b rune) {
	s.Push(builtins.Byte(b))
}

func (s *Stack) PushR(r float64) {
	s.Push(builtins.Real(r))
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