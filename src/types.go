package yasm

import "unsafe"

type (
	u8 = uint8
	i8 = int8

	u16 = uint16
	i16 = int16

	u32 = uint32
	i32 = int32

	u64 = uint64
	i64 = int64

	f32 = float32
	f64 = float64
)

const (
	U8 = iota
	I8
	U16
	I16
	U32
	I32
	U64
	I64
	F32
	F64
	Ptr
)

var (
	__u8__  u8
	__u16__ u16
	__u32__ u32
	__u64__ u64
	__f32__ f32
	__f64__ f64
	__ptr__ uintptr

	SizeOf = map[int]uintptr{
		U8:  unsafe.Sizeof(__u8__),
		I8:  unsafe.Sizeof(__u8__),
		U16: unsafe.Sizeof(__u16__),
		I16: unsafe.Sizeof(__u16__),
		U32: unsafe.Sizeof(__u32__),
		I32: unsafe.Sizeof(__u32__),
		U64: unsafe.Sizeof(__u64__),
		I64: unsafe.Sizeof(__u64__),
		F32: unsafe.Sizeof(__f32__),
		F64: unsafe.Sizeof(__f64__),
		Ptr: unsafe.Sizeof(__ptr__),
	}
)
