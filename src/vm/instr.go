package vm

import "fmt"

var hex = func(x uint32) string {
	return fmt.Sprintf("0x%x", x)
}

var bin = func(x uint32) string {
	return fmt.Sprintf("0b%032b", x)
}

var _str = func(x uint32) string {
	return fmt.Sprintf("%d", x)
}

const (
	tMask = 0xc0000000
	oMask = 0x3f000000
	aMask = 0x00fc0000
	bMask = 0x0003f000
	cMask = 0x00000fc0
	eMask = 0x0000003f
)

// -t---op--- --ra-----rb-- ---rc----ex---
// ttoo oooo aaaa aabb bbbb cccc ccee eeee
func gen(mode uint32, op byte, arg1, arg2, arg3, arg4 uint32) uint32 {
	var i uint32

	i |= mode << 30       // mode
	i |= uint32(op) << 24 // op
	i |= arg1 << 18       // ra
	i |= arg2 << 12       // rb
	i |= arg3 << 6        // rc
	i |= arg4 & eMask     // extra

	return i
}

func read(n, m word) string {
	i := n.UInt32(0)
	tt := (i & tMask) >> 30
	op := (i & oMask) >> 24

	s := ""

	// tt = 0x00 -> ordinary opcode
	// tt = 0x01 -> double word opcode
	// tt = 0x02 -> Integer
	// tt = 0x03 -> Float
	switch tt {
	case 0x01:
		s += " "
		s += opName[byte(op)]
		s += "\t"
		s += fmt.Sprintf("%d", m.Int32(0))
		return s
	case 0x2:
		s += "i"
	case 0x3:
		s += "f"
	default:
		s += " "
	}

	s += opName[byte(op)]
	s += "\t"

	ra := (i & aMask) >> 20
	rb := (i & bMask) >> 14
	rc := (i & cMask) >> 8
	ex := (i & eMask)

	s += regName(ra) + ", "
	s += regName(rb) + ", "
	s += regName(rc) + "    "
	s += hex(ex)

	return s
}
