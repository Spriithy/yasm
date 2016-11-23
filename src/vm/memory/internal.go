package memory

import "fmt"

const (
	// InitialMemorySize dictates the Initial Memory Size
	// for the VM
	InitialMemorySize = 1 << 10
)

var offset = [...]uint64{
	0: 0xFF << 56,
	1: 0xFF << 48,
	2: 0xFF << 40,
	3: 0xFF << 32,
	4: 0xFF << 24,
	5: 0xFF << 16,
	6: 0xFF << 8,
	7: 0xFF << 0,
}

var shift = [...]uint64{
	0: 56,
	1: 48,
	2: 40,
	3: 32,
	4: 24,
	5: 16,
	6: 8,
	7: 0,
}

// Memory is a Double word storage place where every bit of data has a purpose
// Types stored range from
//  -> int8 / uint8
//  -> int64 / uint64
//  -> float64
//  -> strings
//  -> structures
type Memory []uint64

// Init creates a new memory instance to allow users's computations
// and store data
func Init() Memory {
	return make(Memory, InitialMemorySize)
}

var max = func(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (m Memory) Dump(start, end int) {
	for i := start; i < end; i += 2 {
		if i%8 == 0 {
			fmt.Printf("\n%06X.\t", max(i-4, 0))
		}
		fmt.Printf("%02X%02X ", m.getByte(i), m.getByte(i+1))
	}
}

func (m Memory) getByte(addr int) byte {
	ofs := addr / 8
	byt := addr % 8
	return byte((m[ofs] & offset[byt]) >> shift[byt])
}
