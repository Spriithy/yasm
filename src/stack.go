package yasm

import (
	"fmt"
	"unsafe"
)

const (
	StackAlign = unsafe.Sizeof(__u32__)
	StackSize  = 1 << 16
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

func (c *cpu) peek8() u8        { return getu8(c.sp) }
func (c *cpu) peek16() u16      { return getu16(c.sp) }
func (c *cpu) peek32() u32      { return getu32(c.sp) }
func (c *cpu) peek64() u64      { return getu64(c.sp) }
func (c *cpu) peekf32() f32     { return getf32(c.sp) }
func (c *cpu) peekf64() f64     { return getf64(c.sp) }
func (c *cpu) peekPtr() uintptr { return getPtr(c.sp) }

func (c *cpu) pushu8(v u8)   { c.trapOverflow(StackAlign); c.sp -= StackAlign; setu8(c.sp, v) }
func (c *cpu) pushi8(v i8)   { c.trapOverflow(StackAlign); c.sp -= StackAlign; seti8(c.sp, v) }
func (c *cpu) pushu16(v u16) { c.trapOverflow(StackAlign); c.sp -= StackAlign; setu16(c.sp, v) }
func (c *cpu) pushi16(v i16) { c.trapOverflow(StackAlign); c.sp -= StackAlign; seti16(c.sp, v) }
func (c *cpu) pushu32(v u32) { c.trapOverflow(StackAlign); c.sp -= StackAlign; setu32(c.sp, v) }
func (c *cpu) pushi32(v i32) { c.trapOverflow(StackAlign); c.sp -= StackAlign; seti32(c.sp, v) }
func (c *cpu) pushu64(v u64) { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; setu64(c.sp, v) }
func (c *cpu) pushi64(v i64) { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; seti64(c.sp, v) }
func (c *cpu) pushf32(v f32) { c.trapOverflow(StackAlign); c.sp -= StackAlign; setf32(c.sp, v) }
func (c *cpu) pushf64(v f64) { c.trapOverflow(StackAlign * 2); c.sp -= StackAlign * 2; setf64(c.sp, v) }
func (c *cpu) pushPtr(v uintptr) {
	c.trapOverflow(StackAlign * 2)
	c.sp -= StackAlign * 2
	setPtr(c.sp, v)
}

func (c *cpu) pop8() u8 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return getu8(c.sp - StackAlign)
}

func (c *cpu) pop16() u16 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return getu16(c.sp - StackAlign)
}

func (c *cpu) pop32() u32 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return getu32(c.sp - StackAlign)
}

func (c *cpu) pop64() u64 {
	c.trapUnderflow(StackAlign * 2)
	c.sp += StackAlign * 2
	return getu64(c.sp - StackAlign*2)
}

func (c *cpu) popf32() f32 {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign
	return getf32(c.sp - StackAlign)
}

func (c *cpu) popf64() f64 {
	c.trapUnderflow(StackAlign * 2)
	c.sp += StackAlign * 2
	return getf64(c.sp - StackAlign*2)
}

func (c *cpu) popPtr() uintptr {
	c.trapUnderflow(StackAlign)
	c.sp += StackAlign * 2
	return getPtr(c.sp - StackAlign*2)
}

func isDumpAt(p uintptr, pos uintptr) bool {
	return p-1 >= pos-StackAlign && p-1 < pos
}

func (c *cpu) stackDump() {
	for i := len(c.stack); i > StackSize-128; i -= int(StackAlign) {
		// dump pointer
		dp := uintptr(unsafe.Pointer(&c.stack[i-1]))
		if isDumpAt(c.sp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				fmt.Println()
			}
			fmt.Printf("\x1b[44;1m0x%08x SP \x1b[0m\x1b[34m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if isDumpAt(c.bp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				fmt.Println()
			}
			fmt.Printf("\x1b[41;1m0x%08x BP \x1b[0m\x1b[31m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if isDumpAt(c.fp, dp) && i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				fmt.Println()
			}
			fmt.Printf("\x1b[42;1m0x%08x FP \x1b[0m\x1b[32m\uE0B0\x1b[0m ", dp-StackAlign+1)
		} else if i%int(StackAlign) == 0 {
			if i != len(c.stack) {
				fmt.Println()
			}
			fmt.Printf("0x%08x      ", dp-StackAlign+1)
		}

		for j := 0; j < int(StackAlign); j++ {
			fmt.Printf("0x%02x ", c.stack[i-1-j])
		}

		fmt.Print("\x1b[0m")
	}
	fmt.Println()
}
