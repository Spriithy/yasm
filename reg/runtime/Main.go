package main

import (
	. "github.com/Spriithy/Polaroid/reg/runtime/objects"
	colors "github.com/Spriithy/go-colors"
)

func check(e error) {
	if e != nil {
		println(colors.Red(colors.None, e))
		panic(e)
	}
}

func main() {

	dim := 5

	a, err := CreateArray(dim, "null")
	check(err)

	for i := 0; i < dim; i++ {
		col, err := CreateArray(dim, "int")
		check(err)
		err = col.Write(i, Int(1))
		check(err)
		a.Write(i, col)
	}

	println(a.Class(), a.String())

}
