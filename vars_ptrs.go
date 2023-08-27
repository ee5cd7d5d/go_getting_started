package main

import "fmt"

func vars_ptrs() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	d, e := real(c), imag(c)
	fmt.Println(d, e)

	var firstname *string = new(string)
	fmt.Println(firstname)

	*firstname = "Arthur"

	println(firstname)
	println(*firstname)

	another_firstname := "Arthur2"
	fmt.Println(another_firstname)

	ptr := &another_firstname
	fmt.Println(ptr, *ptr)
	another_firstname = "Arthur3"
	fmt.Println(ptr, *ptr)
}
