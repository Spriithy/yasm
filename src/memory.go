package yasm

import "unsafe"

func get8(p uintptr) i8        { return *(*i8)(unsafe.Pointer(p)) }
func get8u(p uintptr) u8       { return *(*u8)(unsafe.Pointer(p)) }
func get16(p uintptr) i16      { return *(*i16)(unsafe.Pointer(p)) }
func get16u(p uintptr) u16     { return *(*u16)(unsafe.Pointer(p)) }
func get32(p uintptr) i32      { return *(*i32)(unsafe.Pointer(p)) }
func get32u(p uintptr) u32     { return *(*u32)(unsafe.Pointer(p)) }
func get64(p uintptr) i64      { return *(*i64)(unsafe.Pointer(p)) }
func get64u(p uintptr) u64     { return *(*u64)(unsafe.Pointer(p)) }
func get32f(p uintptr) f32     { return *(*f32)(unsafe.Pointer(p)) }
func get64f(p uintptr) f64     { return *(*f64)(unsafe.Pointer(p)) }
func getPtr(p uintptr) uintptr { return *(*uintptr)(unsafe.Pointer(p)) }

func set8(p uintptr, v i8)        { *(*i8)(unsafe.Pointer(p)) = v }
func set8u(p uintptr, v u8)       { *(*u8)(unsafe.Pointer(p)) = v }
func set16(p uintptr, v i16)      { *(*i16)(unsafe.Pointer(p)) = v }
func set16u(p uintptr, v u16)     { *(*u16)(unsafe.Pointer(p)) = v }
func set32(p uintptr, v i32)      { *(*i32)(unsafe.Pointer(p)) = v }
func set32u(p uintptr, v u32)     { *(*u32)(unsafe.Pointer(p)) = v }
func set64(p uintptr, v i64)      { *(*i64)(unsafe.Pointer(p)) = v }
func set64u(p uintptr, v u64)     { *(*u64)(unsafe.Pointer(p)) = v }
func set32f(p uintptr, v f32)     { *(*f32)(unsafe.Pointer(p)) = v }
func set64f(p uintptr, v f64)     { *(*f64)(unsafe.Pointer(p)) = v }
func setPtr(p uintptr, v uintptr) { *(*uintptr)(unsafe.Pointer(p)) = v }

func deref8(p []u8, offset uintptr) i8    { return i8(p[offset]) }
func deref8u(p []u8, offset uintptr) u8   { return p[offset] }
func deref16(p []u8, offset uintptr) i16  { return *(*i16)(unsafe.Pointer(&p[offset])) }
func deref16u(p []u8, offset uintptr) u16 { return *(*u16)(unsafe.Pointer(&p[offset])) }
func deref32(p []u8, offset uintptr) i32  { return *(*i32)(unsafe.Pointer(&p[offset])) }
func deref32u(p []u8, offset uintptr) u32 { return *(*u32)(unsafe.Pointer(&p[offset])) }
func deref32f(p []u8, offset uintptr) f32 { return *(*f32)(unsafe.Pointer(&p[offset])) }
func deref64(p []u8, offset uintptr) i64  { return *(*i64)(unsafe.Pointer(&p[offset])) }
func deref64u(p []u8, offset uintptr) u64 { return *(*u64)(unsafe.Pointer(&p[offset])) }
func deref64f(p []u8, offset uintptr) f64 { return *(*f64)(unsafe.Pointer(&p[offset])) }

func (c *cpu) getLocal8(offset int) i8        { return get8(c.fp + uintptr(offset)) }
func (c *cpu) getLocal8u(offset int) u8       { return get8u(c.fp + uintptr(offset)) }
func (c *cpu) getLocal16(offset int) i16      { return get16(c.fp + uintptr(offset)) }
func (c *cpu) getLocal16u(offset int) u16     { return get16u(c.fp + uintptr(offset)) }
func (c *cpu) getLocal32(offset int) i32      { return get32(c.fp + uintptr(offset)) }
func (c *cpu) getLocal32u(offset int) u32     { return get32u(c.fp + uintptr(offset)) }
func (c *cpu) getLocal64(offset int) i64      { return get64(c.fp + uintptr(offset)) }
func (c *cpu) getLocal64u(offset int) u64     { return get64u(c.fp + uintptr(offset)) }
func (c *cpu) getLocal32f(offset int) f32     { return get32f(c.fp + uintptr(offset)) }
func (c *cpu) getLocal64f(offset int) f64     { return get64f(c.fp + uintptr(offset)) }
func (c *cpu) getLocalPtr(offset int) uintptr { return getPtr(c.fp + uintptr(offset)) }

func (c *cpu) setLocal8(offset int, v i8)        { set8(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal8u(offset int, v u8)       { set8u(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal16(offset int, v i16)      { set16(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal16u(offset int, v u16)     { set16u(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal32(offset int, v i32)      { set32(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal32u(offset int, v u32)     { set32u(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal64(offset int, v u64)      { set64u(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal32f(offset int, v f32)     { set32f(c.fp+uintptr(offset), v) }
func (c *cpu) setLocal64f(offset int, v f64)     { set64f(c.fp+uintptr(offset), v) }
func (c *cpu) setLocalPtr(offset int, v uintptr) { setPtr(c.fp+uintptr(offset), v) }

func (c *cpu) pushLocal8(offset int)   { c.push8(get8(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal8u(offset int)  { c.push8u(get8u(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal16(offset int)  { c.push16(get16(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal16u(offset int) { c.push16u(get16u(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal32(offset int)  { c.push32(get32(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal32u(offset int) { c.push32u(get32u(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal64(offset int)  { c.push64(get64(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal64u(offset int) { c.push64u(get64u(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal32f(offset int) { c.push32f(get32f(c.fp + uintptr(offset))) }
func (c *cpu) pushLocal64f(offset int) { c.push64f(get64f(c.fp + uintptr(offset))) }
func (c *cpu) pushLocalPtr(offset int) { c.pushPtr(getPtr(c.fp + uintptr(offset))) }

func (c *cpu) popLocal8(offset int)   { set8(c.fp+uintptr(offset), c.pop8()) }
func (c *cpu) popLocal8u(offset int)  { set8u(c.fp+uintptr(offset), c.pop8u()) }
func (c *cpu) popLocal16(offset int)  { set16(c.fp+uintptr(offset), c.pop16()) }
func (c *cpu) popLocal16u(offset int) { set16u(c.fp+uintptr(offset), c.pop16u()) }
func (c *cpu) popLocal32(offset int)  { set32(c.fp+uintptr(offset), c.pop32()) }
func (c *cpu) popLocal32u(offset int) { set32u(c.fp+uintptr(offset), c.pop32u()) }
func (c *cpu) popLocal64(offset int)  { set64(c.fp+uintptr(offset), c.pop64()) }
func (c *cpu) popLocal64u(offset int) { set64u(c.fp+uintptr(offset), c.pop64u()) }
func (c *cpu) popLocalf32(offset int) { set32f(c.fp+uintptr(offset), c.pop32f()) }
func (c *cpu) popLocalf64(offset int) { set64f(c.fp+uintptr(offset), c.pop64f()) }
func (c *cpu) popLocalPtr(offset int) { setPtr(c.fp+uintptr(offset), c.popPtr()) }

func (c *cpu) getGlobal8(offset int) i8        { return get8(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal8u(offset int) u8       { return get8u(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal16(offset int) i16      { return get16(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal16u(offset int) u16     { return get16u(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal32(offset int) i32      { return get32(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal32u(offset int) u32     { return get32u(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal64(offset int) i64      { return get64(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal64u(offset int) u64     { return get64u(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal32f(offset int) f32     { return get32f(c.bp + uintptr(offset)) }
func (c *cpu) getGlobal64f(offset int) f64     { return get64f(c.bp + uintptr(offset)) }
func (c *cpu) getGlobalPtr(offset int) uintptr { return getPtr(c.bp + uintptr(offset)) }

func (c *cpu) setGlobal8(offset int, v i8)        { set8(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal8u(offset int, v u8)       { set8u(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal16(offset int, v i16)      { set16(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal16u(offset int, v u16)     { set16u(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal32(offset int, v i32)      { set32(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal32u(offset int, v u32)     { set32u(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal64(offset int, v i64)      { set64(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal64u(offset int, v u64)     { set64u(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal32f(offset int, v f32)     { set32f(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobal64f(offset int, v f64)     { set64f(c.bp+uintptr(offset), v) }
func (c *cpu) setGlobalPtr(offset int, v uintptr) { setPtr(c.bp+uintptr(offset), v) }

func (c *cpu) pushGlobal8(offset int)   { c.push8(get8(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal8u(offset int)  { c.push8u(get8u(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal16(offset int)  { c.push16(get16(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal16u(offset int) { c.push16u(get16u(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal32(offset int)  { c.push32(get32(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal32u(offset int) { c.push32u(get32u(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal64(offset int)  { c.push64(get64(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal64u(offset int) { c.push64u(get64u(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal32f(offset int) { c.push32f(get32f(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobal64f(offset int) { c.push64f(get64f(c.bp + uintptr(offset))) }
func (c *cpu) pushGlobalPtr(offset int) { c.pushPtr(getPtr(c.bp + uintptr(offset))) }

func (c *cpu) popGlobal8(offset int)   { set32(c.bp+uintptr(offset), i32(c.pop8())) }
func (c *cpu) popGlobal8u(offset int)  { set32u(c.bp+uintptr(offset), u32(c.pop8u())) }
func (c *cpu) popGlobal16(offset int)  { set32(c.bp+uintptr(offset), i32(c.pop16())) }
func (c *cpu) popGlobal16u(offset int) { set32u(c.bp+uintptr(offset), u32(c.pop16u())) }
func (c *cpu) popGlobal32(offset int)  { set32(c.bp+uintptr(offset), c.pop32()) }
func (c *cpu) popGlobal32u(offset int) { set32u(c.bp+uintptr(offset), c.pop32u()) }
func (c *cpu) popGlobal64(offset int)  { set64(c.bp+uintptr(offset), c.pop64()) }
func (c *cpu) popGlobal64u(offset int) { set64u(c.bp+uintptr(offset), c.pop64u()) }
func (c *cpu) popGlobal32f(offset int) { set32f(c.bp+uintptr(offset), c.pop32f()) }
func (c *cpu) popGlobal64f(offset int) { set64f(c.bp+uintptr(offset), c.pop64f()) }
func (c *cpu) popGlobalPtr(offset int) { setPtr(c.bp+uintptr(offset), c.popPtr()) }
