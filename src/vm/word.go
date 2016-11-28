package small

import (
	"unsafe"
)

type word uint64

// --- GETTERS --- --- --- --- --- --- ---

func (w word) uint32() uint32 {
	return (uint32)(w)
}

func (w word) int32() int32 {
	return (int32)(w)
}

func (w word) uint64() uint64 {
	return (uint64)(w)
}

func (w word) int64() int64 {
	return (int64)(w)
}

func (w word) float32() float32 {
	return *(*float32)(unsafe.Pointer(&w))
}

func (w word) float64() float64 {
	return *(*float64)(unsafe.Pointer(&w))
}

func (w word) bool() bool {
	return w != 0
}

func (w word) char() byte {
	return (byte)(w)
}

// --- SETTER  --- --- --- --- --- --- ---

func (w word) set(x interface{}) {
	switch x.(type) {
	case *word:
		w = *x.(*word)
	case word:
		w = x.(word)
	case uint32:
		w = word(x.(uint32))
	case int32:
		w = word(x.(int32))
	case uint64:
		w = word(x.(uint64))
	case int64:
		w = word(x.(int64))
	case float32:
		w = word(x.(float32))
	case float64:
		w = word(x.(float64))
	case byte:
		w = word(x.(byte))
	case bool:
		switch {
		case x.(bool):
			w = 1
		default:
			w = 0
		}
	}
}
