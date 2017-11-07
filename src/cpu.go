package yasm

type cpu struct {
	stack [StackSize]u8
	sp    uintptr
	fp    uintptr
	bp    uintptr

	ta uintptr // Trap address
	tm string  // Trap message
	ec int     // exit code

	module *Module
	fn     *Function
}
