package main

import "fmt"

func loops() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}

	i = 0
	for i < 5 {
		if i == 3 {
			break
		}
		fmt.Println(i)
		i++
	}

	i = 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}

	i = 0
	for {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i++
	}

	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, v := range slice { // also works with maps
		fmt.Println(i, v)
	}
}
