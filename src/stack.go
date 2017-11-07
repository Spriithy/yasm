package yasm

const StackSize = 1024

var StackAlign = sizeOf[U32]

func (c *cpu) frameSize() int { return int(c.fp - c.sp) }
func (c *cpu) stackSize() int { return int(c.bp - c.sp) }

func (c *cpu) checkOverflow() {
	if c.sp-StackAlign < c.bp-StackSize {
		c.ta = c.sp - StackAlign
		c.trap(StackOverflow)
	}
}

func (c *cpu) checkUnderflow() {
	if c.sp+StackAlign > c.bp {
		c.ta = c.sp + StackAlign
		c.trap(StackUnderflow)
	}
}

func (c *cpu) pushu8(v u8)     { c.checkOverflow(); c.sp -= StackAlign; setu8(c.sp, v) }
func (c *cpu) pushi8(v i8)     { c.checkOverflow(); c.sp -= StackAlign; seti8(c.sp, v) }
func (c *cpu) pushu16(v u16)   { c.checkOverflow(); c.sp -= StackAlign; setu16(c.sp, v) }
func (c *cpu) pushi16(v i16)   { c.checkOverflow(); c.sp -= StackAlign; seti16(c.sp, v) }
func (c *cpu) pushu32(v u32)   { c.checkOverflow(); c.sp -= StackAlign; setu32(c.sp, v) }
func (c *cpu) pushi32(v i32)   { c.checkOverflow(); c.sp -= StackAlign; seti32(c.sp, v) }
func (c *cpu) pushu64(v u64)   { c.checkOverflow(); c.sp -= StackAlign * 2; setu64(c.sp, v) }
func (c *cpu) pushi64(v i64)   { c.checkOverflow(); c.sp -= StackAlign * 2; seti64(c.sp, v) }
func (c *cpu) pushf32(v f32)   { c.checkOverflow(); c.sp -= StackAlign; setf32(c.sp, v) }
func (c *cpu) pushf64(v f64)   { c.checkOverflow(); c.sp -= StackAlign * 2; setf64(c.sp, v) }
func (c *cpu) pushAddr(v addr) { c.checkOverflow(); c.sp -= StackAlign * 2; setAddr(c.sp, v) }

func (c *cpu) pop8() u8 {
	c.checkUnderflow()
	c.sp += StackAlign
	return getu8(c.sp - StackAlign)
}

func (c *cpu) pop16() u16 {
	c.checkUnderflow()
	c.sp += StackAlign
	return getu16(c.sp - StackAlign)
}

func (c *cpu) pop32() u32 {
	c.checkUnderflow()
	c.sp += StackAlign
	return getu32(c.sp - StackAlign)
}

func (c *cpu) pop64() u64 {
	c.checkUnderflow()
	c.sp += StackAlign * 2
	return getu64(c.sp - StackAlign*2)
}

func (c *cpu) popf32() f32 {
	c.checkUnderflow()
	c.sp += StackAlign
	return getf32(c.sp - StackAlign)
}

func (c *cpu) popf64() f64 {
	c.checkUnderflow()
	c.sp += StackAlign * 2
	return getf64(c.sp - StackAlign*2)
}
func (c *cpu) popAddr() addr {
	c.checkUnderflow()
	c.sp += StackAlign * 2
	return getAddr(c.sp - StackAlign*2)
}
