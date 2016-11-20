package main

import . "github.com/Spriithy/Polaroid/src/builtins"

func main() {
	a := CreateArray(10, "int")
	println(a.Class(), a.String())

	t := CreateTable()
	println(t.Class(), t.String())
	t.Set("array", a)
	a.Set(3, Int(8))
	println(t.Class(), t.String())
	println(a.Class(), a.String())

	e := Enum("Foo", "Bar", 3)
	println(e.Class(), e.String())
}
