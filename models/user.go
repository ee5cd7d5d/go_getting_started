package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // By using pointers the data will be global
	nextID = 1
)
