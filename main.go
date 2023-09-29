package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ee5cd7d5d/go_getting_started/controllers"
)

func main() {
	//getting_started_lessons()
	//port := 3000
	//_, err := startWebServer(port, 2)
	//fmt.Println(err)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting...")
	// Stuff
	fmt.Println("Started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, errors.New("bad error")
}

func getting_started_lessons() {
	//vars_ptrs()
	//consts()
	//collections()
	//loops()
	switches()
}
