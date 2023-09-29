package main

import "fmt"

type HTTPRequest struct {
	Method string
}

func switches() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		fmt.Println("Get request")
		fallthrough
	case "TEAPOT":
		fmt.Println("Teapot request")
	default:
		fmt.Println("What did you doooo")
	}
}
