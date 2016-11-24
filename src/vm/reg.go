package vm

import "github.com/Spriithy/go-colors"

func regName(r uint32) string {
	str := _str // Alias

	switch {
	// Reserved cases
	case r == 0: // Zero register
		return "#0"
	case r == 1:
		return "ra"
	case r == 54, r == 55:
		return colors.Red(colors.None, "r"+str(r))
	case r == 56:
		return "rx"
	case r == 57: // Program Counter
		return "pc"
	case r == 58: // Interrupt address
		return "ia"
	case r == 59: // Interrupt message
		return "im"
	case r == 60: // Interrupt return
		return "ir"
	case r == 61: // Flags
		return "fl"
	case r == 62, r == 63:
		return "s" + str(r)

	// General purposes registers
	case r > 0 && r <= 16: // 16 Function arguments
		return "a" + str(r-1)
	case r > 16 && r <= 48: // 32 Temporaries
		return "r" + str(r-17)
	case r > 48 && r <= 54: // 5 Return values
		return "v" + str(r-49)
	}

	return "r" + str(r)
}
