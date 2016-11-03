package main

import "github.com/Spriithy/Polaroid/src/runtime"

func main() {

	vm := runtime.VirtualMachine()
	println(vm.StackString())

}
