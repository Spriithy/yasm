package vm

import (
	"os"

	"fmt"

	"unsafe"

	colors "github.com/Spriithy/go-colors"
)

func assert(cond bool, msg string) {
	if !cond {
		println(">", colors.Red(colors.None, msg+" error"))
		os.Exit(1)
	}
}

// Type denotes the type held by some data
type Type uint8

const (
	// NullType means the data is both untyped and null
	NullType = Type(iota)

	// ByteType represents data with 8 bits
	ByteType

	// WordType is the basic unit of 16 bits
	WordType

	// DoubleWordType is twice a word -> 32 bits
	DoubleWordType

	// QuadWordType is twic a DoubleWord -> 64 bits
	QuadWordType
)

// A DataUnit represents either a
//  Byte
//  Word
//  DoubleWord
//  QuadWord
type DataUnit interface {
	Type() Type
}

// A Byte is a piece of 8 unsigned bits
type Byte uint8

// Type returns the ByteType constant
func (b Byte) Type() Type {
	return ByteType
}

// Word is used to shift the byte into either first or second place
func (b Byte) Word(place uint) *Word {
	assert(place <= 1, "word overflow")
	x := new(Word)
	*x = Word(b) << (place * 8)
	return x
}

// DoubleWord is used to shift the byte into any place in a double-word
func (b Byte) DoubleWord(place uint) *DoubleWord {
	assert(place <= 3, "double-word overflow")
	x := new(DoubleWord)
	*x = DoubleWord(b) << (place * 8)
	return x
}

// QuadWord is used to shift a byte into any place in a quad-word
func (b Byte) QuadWord(place uint) *QuadWord {
	assert(place <= 7, "quad-word overflow")
	x := new(QuadWord)
	*x = QuadWord(b) << (place * 8)
	return x
}

// Int8 returns an 8-bits signed integer from the Byte
func (b Byte) Int8() int8 {
	var r int8
	r |= int8(b) << 0
	return r
}

// UInt8 returns the 8-bits unsigned integer from the Byte
func (b Byte) UInt8() uint8 {
	var r uint8
	r |= uint8(b) << 0
	return r
}

func (b Byte) String() string {
	return fmt.Sprintf("%02X", uint8(b))
}

// A Word is two Bytes and the basic architecure size
type Word uint16

// Byte returns the n-th byte from the word
func (w Word) Byte(place uint) *Byte {
	assert(place <= 1, "word overflow")
	x := new(Byte)
	*x = Byte(w >> (place * 8))
	return x
}

// DoubleWord returns the Word shifted into a new DoubleWord
func (w Word) DoubleWord(place uint) *DoubleWord {
	assert(place <= 1, "double-word overflow")
	x := new(DoubleWord)
	*x = DoubleWord(w) << (place * 16)
	return x
}

// QuadWord returns the Word shifted into a new QuadWord
func (w Word) QuadWord(place uint) *QuadWord {
	assert(place <= 3, "quad-word overflow")
	x := new(QuadWord)
	*x = QuadWord(w) << (place * 16)
	return x
}

// Int16 returns a 16-bits signed integer from the Word
func (w Word) Int16() int16 {
	var r int16
	r |= int16(*w.Byte(0)) << 0
	r |= int16(*w.Byte(1)) << 8
	return r
}

// UInt16 returns the 16-bits unsigned integer from the Word
func (w Word) UInt16() uint16 {
	var r uint16
	r |= uint16(*w.Byte(0)) << 0
	r |= uint16(*w.Byte(1)) << 8
	return r
}

func (w Word) String() string {
	return fmt.Sprintf("%04X", uint16(w))
}

// --- DOUBLEWORD --- --- --- --- --- --- --- --- --- --- --- --- ---

// A DoubleWord is two Word units groups
type DoubleWord uint32

// Byte returns the n-th byte from the DoubleWord
func (dw DoubleWord) Byte(place uint) *Byte {
	assert(place <= 3, "double-word overflow")
	x := new(Byte)
	*x = Byte(dw >> (place * 8))
	return x
}

// Word returns the n-th word from the DoubleWord
func (dw DoubleWord) Word(place uint) *Word {
	assert(place <= 1, "double-word overflow")
	x := new(Word)
	*x = Word(dw >> (place * 16))
	return x
}

// QuadWord returns the Word shifted into a new QuadWord
func (dw DoubleWord) QuadWord(place uint) *QuadWord {
	assert(place <= 1, "quad-word overflow")
	x := new(QuadWord)
	*x = QuadWord(dw) << (place * 32)
	return x
}

// Float32 returns a 32-bits IEEE floating point number from the 32-bits
// of the DoubleWord
func (dw DoubleWord) Float32() float32 {
	bits := dw.UInt32()
	return *(*float32)(unsafe.Pointer(&bits))
}

// Int32 returns a 32-bits signed integer from the DoubleWord
func (dw DoubleWord) Int32() int32 {
	var r int32
	r |= int32(*dw.Byte(0)) << 0
	r |= int32(*dw.Byte(1)) << 8
	r |= int32(*dw.Byte(2)) << 16
	r |= int32(*dw.Byte(3)) << 24
	return r
}

// UInt32 returns the 32-bits signed integer from the DoubleWord
func (dw DoubleWord) UInt32() uint32 {
	var r uint32
	r |= uint32(*dw.Byte(0)) << 0
	r |= uint32(*dw.Byte(1)) << 8
	r |= uint32(*dw.Byte(2)) << 16
	r |= uint32(*dw.Byte(3)) << 24
	return r
}

func (dw DoubleWord) String() string {
	return fmt.Sprintf("%08X", uint32(dw))
}

// --- QUADWORD --- --- --- --- --- --- --- --- --- --- --- --- ---

// A QuadWord is four Word units grouped
type QuadWord uint64

// Byte returns the n-th byte from the QuadWord
func (qw QuadWord) Byte(place uint) *Byte {
	assert(place <= 7, "quad-word overflow")
	x := new(Byte)
	*x = Byte(qw >> (place * 8))
	return x
}

// Word returns the n-th word from the QuadWord
func (qw QuadWord) Word(place uint) *Word {
	assert(place <= 3, "quad-word overflow")
	x := new(Word)
	*x = Word(qw >> (place * 16))
	return x
}

// DoubleWord returns the n-th DoubleWord from the QuadWord
func (qw QuadWord) DoubleWord(place uint) *DoubleWord {
	assert(place <= 7, "quad-word overflow")
	x := new(DoubleWord)
	*x = DoubleWord(qw >> (place * 32))
	return x
}

// Int64 returns a 64-bits signed integer from the quadword
func (qw QuadWord) Int64() int64 {
	var r int64
	r |= int64(*qw.Byte(0)) << 0
	r |= int64(*qw.Byte(1)) << 8
	r |= int64(*qw.Byte(2)) << 16
	r |= int64(*qw.Byte(3)) << 24
	r |= int64(*qw.Byte(4)) << 32
	r |= int64(*qw.Byte(5)) << 40
	r |= int64(*qw.Byte(6)) << 48
	r |= int64(*qw.Byte(7)) << 56
	return r
}

// UInt64 returns the 64-bits signed integer from the quadword
func (qw QuadWord) UInt64() uint64 {
	var r uint64
	r |= uint64(*qw.Byte(0)) << 0
	r |= uint64(*qw.Byte(1)) << 8
	r |= uint64(*qw.Byte(2)) << 16
	r |= uint64(*qw.Byte(3)) << 24
	r |= uint64(*qw.Byte(4)) << 32
	r |= uint64(*qw.Byte(5)) << 40
	r |= uint64(*qw.Byte(6)) << 48
	r |= uint64(*qw.Byte(7)) << 56
	return r
}

// Float32 returns a 32-bits IEEE floating point number from the 32 last
// or first bits of the QuadWord
//func (qw QuadWord) Float32(place uint) float32 {
//	bits := qw.DoubleWord(place).UInt32() // Range check happens here
//	return *(*float32)(unsafe.Pointer(&bits))
//}

// Float64 returns a 64-bits IEEE floating point number
func (qw QuadWord) Float64() float64 {
	bits := qw.UInt64()
	return *(*float64)(unsafe.Pointer(&bits))
}

func (qw QuadWord) String() string {
	return fmt.Sprintf("%016X", uint64(qw))
}
