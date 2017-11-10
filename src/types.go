package yasm

import "unsafe"

type (
	i8 = int8
	u8 = uint8

	i16 = int16
	u16 = uint16

	i32 = int32
	u32 = uint32

	i64 = int64
	u64 = uint64

	f32 = float32
	f64 = float64
)

const (
	I8 = iota
	U8
	I16
	U16
	I32
	U32
	I64
	U64
	F32
	F64
	Ptr
	Array = 1 << 7
)

var (
	__u8__  u8
	__u16__ u16
	__u32__ u32
	__u64__ u64
	__f32__ f32
	__f64__ f64
	__ptr__ uintptr

	SizeOf = []uintptr{
		unsafe.Sizeof(__u8__),
		unsafe.Sizeof(__u8__),
		unsafe.Sizeof(__u16__),
		unsafe.Sizeof(__u16__),
		unsafe.Sizeof(__u32__),
		unsafe.Sizeof(__u32__),
		unsafe.Sizeof(__u64__),
		unsafe.Sizeof(__u64__),
		unsafe.Sizeof(__f32__),
		unsafe.Sizeof(__f64__),
		unsafe.Sizeof(__ptr__),
	}

	TypeName = []string{
		"i8",
		"u8",
		"i16",
		"u16",
		"i32",
		"u32",
		"i64",
		"u64",
		"f32",
		"f64",
		"ptr",
	}
)
