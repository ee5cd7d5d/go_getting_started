package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	// this is taking the interface we defined in the other file
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc) // this is needed because the ending slash makes it a different address
}
