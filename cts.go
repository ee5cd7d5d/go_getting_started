package main

import "fmt"

const (
	first = iota + 6
	first_bis
	second = 2 << iota
	third
)

const (
	fourth = iota
)

func consts() {
	const pi = 3.141509
	fmt.Println(pi)

	// Implicitly typed constants
	const c = 3
	fmt.Println(c + 3)
	fmt.Println(c + 2.1)
	// Explicit
	const d int = 3
	fmt.Println(d + 3)
	fmt.Println(float32(d) + 2.1)

	fmt.Println(first)
	fmt.Println(first_bis)
	fmt.Println(second) // iota is like a global iterator, increasing with every use
	fmt.Println(third)
	fmt.Println(fourth) // This belongs to another constant block so it resets
}
