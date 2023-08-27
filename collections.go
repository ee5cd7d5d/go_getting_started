package main

import "fmt"

func collections() {
	// Arrays
	// Arrays are fixed length
	// Long syntax
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])
	// fmt.Println(arr[4]) // compile error (out of bounds)

	// Short syntax
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Slices
	slice := arr[:]
	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27 // Slice builds on top of the array
	fmt.Println(arr, slice)

	// Initializing as a slice directly
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	// Appending
	slice2 = append(slice2, 4, 42, 77)
	fmt.Println(slice2)
	// Slice the slice
	s3 := slice2[1:]
	s4 := slice2[:2]
	s5 := slice2[1:2]
	fmt.Println(s3, s4, s5)

	// Maps
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])
	m["foo"] = 27
	fmt.Println(m)
	delete(m, "foo")
	fmt.Println(m)

	// Struct
	type user struct {
		ID        int
		FirstName string
		LastName  string
	} // This is in scope of main! Could be moved out of scope
	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Doyle"
	fmt.Println(u)
	fmt.Println(u.LastName)

	u2 := user{
		ID:        1,
		FirstName: "Arthur",
		LastName:  "Doyle",
	}
	fmt.Println(u2)
}
