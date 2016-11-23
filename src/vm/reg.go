package vm

import "github.com/Spriithy/go-colors"

func regName(r uint32) string {
	switch {
	// Reserved cases
	case r == 0: // Zero register
		return "#0"
	case r == 1:
		return "xra"
	case r == 53, r == 54:
		return colors.Red(colors.None, "r"+str(r))
	case r == 55: // Program Counter
		return "pc"
	case r == 56: // Frame pointer
		return "fp"
	case r == 57: // Stack pointer
		return "esp"
	case r == 58: // Interrupt address
		return "ia"
	case r == 59: // Interrupt message
		return "im"
	case r == 60: // Interrupt return
		return "ir"
	case r == 61: // Flags
		return "fl"
	case r == 62, r == 63:
		return colors.Red(colors.None, "r"+str(r))

	// General purposes registers
	case r > 0 && r <= 16: // 16 Function arguments
		return "a" + str(r-1)
	case r > 16 && r <= 48: // 32 Temporaries
		return "r" + str(r-17)
	case r > 48 && r <= 52: // 3 Return values
		return "v" + str(r-49)
	}

	return "r" + str(r)
}
