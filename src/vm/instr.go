package vm

import "fmt"

var hex = func(x uint32) string {
	return fmt.Sprintf("0x%x", x)
}

var bin = func(x uint32) string {
	return fmt.Sprintf("0b%b", x)
}

var str = func(x uint32) string {
	return fmt.Sprintf("%d", x)
}

const (
	bcOpcodeMask = 0xfc000000
	bcAMask      = 0x03f00000
	bcBMask      = 0x000fc000
	bcCMask      = 0x00003f00
	bcXMask      = 0x000000c0
	bcYMask      = 0x0000003f
)

// oooo ooaa aaaa bbbb bbcc cccc xxyy yyyy
func gen(mode uint32, op byte, arg1, arg2, arg3, arg4 uint32) uint32 {
	var i uint32

	i |= uint32(op) << 26
	i |= arg1 << 20
	i |= arg2 << 14
	i |= arg3 << 8
	i |= mode << 6
	i |= arg4 & bcYMask

	return i
}

func read(i uint32) string {
	op := (i & bcOpcodeMask) >> 26
	ax := (i & bcXMask) >> 6

	s := ""
	switch ax {
	case 0x2:
		s += "i"
	case 0x3:
		s += "f"
	default:
		s += " "
	}

	s += opName[byte(op)]
	s += "\t"

	ra := (i & bcAMask) >> 20
	rb := (i & bcBMask) >> 14
	rc := (i & bcCMask) >> 8
	ay := (i & bcYMask)

	s += regName(ra) + ", "
	s += regName(rb) + ", "
	s += regName(rc) + "    "
	s += hex(ay)

	return s
}
