package main

import . "github.com/Spriithy/Polaroid/reg/runtime/objects"

func main() {

	dim := 5

	defa, _ := CreateArray(0, Int(0))
	a, err := CreateArray(dim, defa.Zero())

	if err != nil {
		panic(err)
	}

	for i := 0; i < dim; i++ {
		c, _ := CreateArray(dim, Int(0))
		err = c.Write(i, Int(1))
		if err != nil {
			panic(err)
		}
		err = a.Write(i, c)
		if err != nil {
			panic(err)
		}
	}

	println(a.Class(), a.String())

}
