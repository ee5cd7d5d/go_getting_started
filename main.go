package main

import (
	"fmt"

	"github.com/ee5cd7d5d/go_getting_started/models"
)

func main() {
	//getting_started_lessons()
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}

func getting_started_lessons() {
	vars_ptrs()
	consts()
	collections()
}
