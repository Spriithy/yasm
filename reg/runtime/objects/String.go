package objects

import (
	"errors"
	"fmt"
	"sync"
)

type StringLit struct {
	size int
	sync.RWMutex
	bytes []byte
}

func String(src string) *StringLit {
	return &StringLit{size: len(src), bytes: []byte(src)}
}

func (s *StringLit) Class() string {
	return "string"
}

func (s *StringLit) Copy() Any {
	return &StringLit{size: s.size, bytes: s.bytes}
}

func (s *StringLit) Read(at int) (byte, error) {
	if at >= s.size || at < 0 {
		return 0, errors.New("string index out of bounds")
	}

	s.RLock()
	defer s.RUnlock()

	return s.bytes[at], nil
}

func (s *StringLit) Write(at int, b byte) error {
	if at >= s.size || at < 0 {
		return errors.New("string index out of bounds")
	}

	s.Lock()
	defer s.Unlock()
	s.bytes[at] = b
	return nil
}

func (s *StringLit) String() string {
	s.RLock()
	defer s.RUnlock()
	return fmt.Sprintf("%#v", string(s.bytes))
}
