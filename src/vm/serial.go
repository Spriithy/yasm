package vm

func int32FromBytes(b0, b1, b2, b3 byte) int32 {
	return int32(b0)<<24 | int32(b1)<<16 | int32(b2)<<8 | int32(b3)
}

func uint32FromBytes(b0, b1, b2, b3 byte) uint32 {
	return uint32(b0)<<24 | uint32(b1)<<16 | uint32(b2)<<8 | uint32(b3)
}

func uint32ToBytes(i0 uint32) (b0, b1, b2, b3 byte) {
	return byte((i0 >> 24) & 0xFF), byte((i0 >> 16) & 0xFF), byte((i0 >> 8) & 0xFF), byte(i0 & 0xFF)
}

func int32ToBytes(i0 int32) (b0, b1, b2, b3 byte) {
	return byte((i0 >> 24) & 0xFF), byte((i0 >> 16) & 0xFF), byte((i0 >> 8) & 0xFF), byte(i0 & 0xFF)
}

func int16FromBytes(b0, b1 byte) int16 {
	return int16(b0)<<8 | int16(b1)
}

func uint16FromBytes(b0, b1 byte) uint16 {
	return uint16(b0)<<8 | uint16(b1)
}

func uint16ToBytes(i0 uint16) (byte, byte) {
	return byte(i0 >> 8), byte(i0 & 0x00FF)
}

func int16ToBytes(i0 int16) (byte, byte) {
	return byte(i0 >> 8), byte(i0 & 0x00FF)
}
