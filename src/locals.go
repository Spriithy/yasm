package yasm

func (c *cpu) getLocal8(n int) u8        { return getu8(c.fp + uintptr(n)) }
func (c *cpu) getLocal16(n int) u16      { return getu16(c.fp + uintptr(n)) }
func (c *cpu) getLocal32(n int) u32      { return getu32(c.fp + uintptr(n)) }
func (c *cpu) getLocal64(n int) u64      { return getu64(c.fp + uintptr(n)) }
func (c *cpu) getLocalf32(n int) f32     { return getf32(c.fp + uintptr(n)) }
func (c *cpu) getLocalf64(n int) f64     { return getf64(c.fp + uintptr(n)) }
func (c *cpu) getLocalPtr(n int) uintptr { return getPtr(c.fp + uintptr(n)) }
