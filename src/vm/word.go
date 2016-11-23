package vm

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/Spriithy/go-colors"
)

func assert(cond bool, msg string) {
	if !cond {
		println(">", colors.Red(colors.None, msg+" error"))
		os.Exit(1)
	}
} // --- WORD --- --- --- --- --- --- --- --- --- --- --- --- ---

// A Word is four Word units grouped
type Word uint64

// Int8 returns the n-th int8 in the word
func (w Word) Int8(place uint) int8 {
	assert(place <= 7, "word overflow")
	return int8(w >> (place * 8))
}

// Int16 returns the n-th int16 in the word
func (w Word) Int16(place uint) int16 {
	assert(place <= 3, "word overflow")
	return int16(w >> (place * 16))
}

// Int32 returns the n-th int32 in the word
func (w Word) Int32(place uint) int32 {
	assert(place <= 1, "word overflow")
	return int32(w >> (place * 32))
}

// Int64 returns a 64-bits signed integer from the Word
func (w Word) Int64() int64 {
	return int64(w.Int32(1))<<32 | int64(w.Int32(0))
}

// UInt8 returns the n-th uint8 in the word
func (w Word) UInt8(place uint) uint8 {
	assert(place <= 7, "word overflow")
	return uint8(w >> (place * 8))
}

// UInt16 returns the n-th uint16 in the word
func (w Word) UInt16(place uint) uint16 {
	assert(place <= 3, "word overflow")
	return uint16(w >> (place * 16))
}

// UInt32 returns the n-th uint32 in the word
func (w Word) UInt32(place uint) uint32 {
	assert(place <= 1, "word overflow")
	return uint32(w >> (place * 32))
}

// UInt64 returns the 64-bits signed integer from the Word
func (w Word) UInt64() uint64 {
	return uint64(w)
}

// Float32 returns a 32-bits IEEE floating point number from the 32 last
// or first bits of the Word
func (w Word) Float32(place uint) float32 {
	bits := w.UInt32(place) // Range check happens here
	return *(*float32)(unsafe.Pointer(&bits))
}

// Float64 returns a 64-bits IEEE floating point number
func (w Word) Float64() float64 {
	bits := w.UInt64()
	return *(*float64)(unsafe.Pointer(&bits))
}

func (w Word) String() string {
	return fmt.Sprintf("%016X", uint64(w))
}
