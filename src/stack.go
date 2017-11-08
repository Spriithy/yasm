package yasm

import (
	"fmt"
	"unsafe"
)

const (
	StackAlign   = unsafe.Sizeof(__u32__)
	StackSize    = 1 << 16
	StackMinDump = 128
)

func (c *cpu) frameSize() int { return int(c.fp - c.sp) }
func (c *cpu) stackSize() int { return int(c.bp - c.sp) }

func (c *cpu) trapOverflow(offset uintptr) {
	if c.sp-offset < c.bp-StackSize {
		c.ta = c.sp - offset
		c.Trap(StackOverflow)
	}
}

func (c *cpu) trapUnderflow(offset uintptr) {
	if c.sp+offset > c.bp {
		c.ta = c.sp + offset
		c.Trap(StackUnderflow)
	}
}

func (c *cpu) peek8() i8        { return get8(c.sp) }
func (c *cpu) peek8u() u8       { return get8u(c.sp) }
func (c *cpu) peek16() i16      { return get16(c.sp) }
func (c *cpu) peek16u() u16     { return get16u(c.sp) }
func (c *cpu) peek32() i32      { return get32(c.sp) }
func (c *cpu) peek32u() u32     { return get32u(c.sp) }
func (c *cpu) peek64() i64      { return get64(c.sp) }
func (c *cpu) peek64u() u64     { return get64u(c.sp) }
func (c *cpu) peek32f() f32     { return get32f(c.sp) }
func (c *cpu) peek64f() f64     { return get64f(c.sp) }
func (c *cpu) peekPtr() uintptr { return getPtr(c.sp) }

func (c *cpu) push8(v i8)    { c.push32(i32(v)) }
func (c *cpu) push8u(v u8)   { c.push32u(u32(v)) }
func (c *cpu) push16(v i16)  { c.push32(i32(v)) }
func (c *cpu) push16u(v u16) { c.push32u(u32(v)) }
func (c *cpu) push32(v i32)  { c.trapOverflow(StackAlign); c.sp -= StackAlign; set32(c.sp, v) }
func (c *cpu) push32u(v u32) { c.trapOverflow(StackAlign); c.sp -= StackAlign; set32u(c.sp, v) }
func (c *cpu) push64(v i64)  { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; set64(c.sp, v) }
func (c *cpu) push64u(v u64) { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; set64u(c.sp, v) }
func (c *cpu) push32f(v f32) { c.trapOverflow(StackAlign); c.sp -= StackAlign; set32f(c.sp, v) }
func (c *cpu) push64f(v f64) { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; set64f(c.sp, v) }
func (c *cpu) pushPtr(v uintptr) {
	c.trapOverflow(StackAlign * 2)
	c.sp -= StackAlign * 2
	setPtr(c.sp, v)
}

func (c *cpu) pop8() i8 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get8(c.sp - StackAlign)
}

func (c *cpu) pop8u() u8 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get8u(c.sp - StackAlign)
}

func (c *cpu) pop16() i16 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get16(c.sp - StackAlign)
}

func (c *cpu) pop16u() u16 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get16u(c.sp - StackAlign)
}

func (c *cpu) pop32() i32 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get32(c.sp - StackAlign)
}

func (c *cpu) pop32u() u32 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get32u(c.sp - StackAlign)
}

func (c *cpu) pop64() i64 {
	c.trapUnderflow(StackAlign * 2)
	c.sp += StackAlign * 2
	return get64(c.sp - StackAlign*2)
}

func (c *cpu) pop64u() u64 {
	c.trapUnderflow(StackAlign * 2)
	c.sp += StackAlign * 2
	return get64u(c.sp - StackAlign*2)
}

func (c *cpu) pop32f() f32 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return get32f(c.sp - StackAlign)
}

func (c *cpu) pop64f() f64 {
	c.trapUnderflow(StackAlign * 2)
	c.sp += StackAlign * 2
	return get64f(c.sp - StackAlign*2)
}

func (c *cpu) popPtr() uintptr {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign * 2
	return getPtr(c.sp - StackAlign*2)
}

func isDumpAt(p uintptr, pos uintptr) bool {
	return p-1 >= pos-StackAlign && p-1 < pos
}

func (c *cpu) stackDump(depth int) {
	if depth < StackMinDump {
		depth = StackMinDump
	}

	if depth > c.stackSize() {
		depth = c.stackSize() + 2*int(StackAlign)
	}

	lim := c.sp + uintptr(depth)
	for p := c.sp + StackAlign; p < lim; p += StackAlign {
		// dump pointer
		dp := p - 1 //uintptr(unsafe.Pointer(&c.stack[p-1]))
		if isDumpAt(c.sp, dp) && p%StackAlign == 0 {
			if p != lim {
				fmt.Println()
			}
			fmt.Printf("\x1b[44;1m0x%08x SP \x1b[0m\x1b[34m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if isDumpAt(c.bp, dp) && p%StackAlign == 0 {
			if p != lim {
				fmt.Println()
			}
			fmt.Printf("\x1b[41;1m0x%08x BP \x1b[0m\x1b[31m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if isDumpAt(c.fp, dp) && p%StackAlign == 0 {
			if p != lim {
				fmt.Println()
			}
			fmt.Printf("\x1b[42;1m0x%08x FP \x1b[0m\x1b[32m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if p%StackAlign == 0 {
			if p != lim {
				fmt.Println()
			}
			fmt.Printf("0x%08x      ", dp-StackAlign+1)
		}

		for j := 0; j < int(StackAlign); j++ {
			fmt.Printf("0x%02x ", get8u(p-uintptr(1+j)))
		}

		fmt.Print("\x1b[0m")
	}
	fmt.Println()
}
