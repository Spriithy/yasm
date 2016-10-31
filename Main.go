package main

import "github.com/Spriithy/Polaroid/src/runtime"

func main() {

	L := runtime.NewList(10)
	for i := 0; i < L.Length(); i++ {
		L.Set(i, runtime.Wrap(int32(i * i)))
	}

	println(L.Equals(L))

	L.Append(runtime.Wrap("Goo"))

	L2 := L.SubList(3, 9)
	for i := 0; i < L2.Length(); i++ {
		println(L.Get(i).ToString(), L2.Get(i).ToString())
	}

}
