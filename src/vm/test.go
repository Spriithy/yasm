package vm

import "strconv"

func binFmt(src int64) string {
	str := strconv.FormatInt(src, 2)
	ss := ""
	for i, c := range str {
		if i%4 == 0 && i > 0 {
			ss += " "
		}
		ss += string(c)
	}
	return ss
}

// Main is
func Main() {
	x := int64((2<<(6-1) - 1)) << 25
	println(binFmt(x))
	x = x & 0xFC000000
	println(binFmt(x))
}
