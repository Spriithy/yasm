package main

import "github.com/Spriithy/Polaroid/src/runtime"

func main() {

	vm := runtime.VirtualMachine()
	for i := 0; i < 10; i++ {
		vm.Push(float64(i * i))
	}
	println(vm.StackString())

}
