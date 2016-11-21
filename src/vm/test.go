package vm

// Main is
func Main() {
	var x uint32 = uint32(mov)<<26 | uint32(blt)<<21 | uint32(lb)<<16 | 1673
	decode(x)
}
