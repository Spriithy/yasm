package small

import "unsafe"

func getu8(p addr) u8     { return *(*u8)(unsafe.Pointer(p)) }
func geti8(p addr) i8     { return *(*i8)(unsafe.Pointer(p)) }
func getu16(p addr) u16   { return *(*u16)(unsafe.Pointer(p)) }
func geti16(p addr) i16   { return *(*i16)(unsafe.Pointer(p)) }
func getu32(p addr) u32   { return *(*u32)(unsafe.Pointer(p)) }
func geti32(p addr) i32   { return *(*i32)(unsafe.Pointer(p)) }
func getu64(p addr) u64   { return *(*u64)(unsafe.Pointer(p)) }
func geti64(p addr) i64   { return *(*i64)(unsafe.Pointer(p)) }
func getf32(p addr) f32   { return *(*f32)(unsafe.Pointer(p)) }
func getf64(p addr) f64   { return *(*f64)(unsafe.Pointer(p)) }
func getAddr(p addr) addr { return *(*addr)(unsafe.Pointer(p)) }

func setu8(p addr, v u8)     { *(*u8)(unsafe.Pointer(p)) = v }
func seti8(p addr, v i8)     { *(*i8)(unsafe.Pointer(p)) = v }
func setu16(p addr, v u16)   { *(*u16)(unsafe.Pointer(p)) = v }
func seti16(p addr, v i16)   { *(*i16)(unsafe.Pointer(p)) = v }
func setu32(p addr, v u32)   { *(*u32)(unsafe.Pointer(p)) = v }
func seti32(p addr, v i32)   { *(*i32)(unsafe.Pointer(p)) = v }
func setu64(p addr, v u64)   { *(*u64)(unsafe.Pointer(p)) = v }
func seti64(p addr, v i64)   { *(*i64)(unsafe.Pointer(p)) = v }
func setf32(p addr, v f32)   { *(*f32)(unsafe.Pointer(p)) = v }
func setf64(p addr, v f64)   { *(*f64)(unsafe.Pointer(p)) = v }
func setAddr(p addr, v addr) { *(*addr)(unsafe.Pointer(p)) = v }
