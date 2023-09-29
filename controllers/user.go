package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

// this is implementing the interface at golang.org/pkg/http/#Handler
// and the compiler will automatically see what we are doing
// this is a method, so I am connecting this to the userController struct
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user controller"))
}

// constructor function is the golang way to implement constructor methods
// the naming new<bla> is the convetion for constructor functions
// returning pointers is also a convetion to avoid copy operations
func newUserController() *userController {
	return &userController{ // this returns the address of the var (the pointer) and we dont have to worry about the scope, go takes care of that for us
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
