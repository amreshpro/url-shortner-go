package controller

import (
	"net/http"
)

type User struct {
	id int
	name  string
	email string
	age   int16
	sex   string
}

func (u User)getUser(w http.ResponseWriter, r http.Request) User {
	return User{
		1,
		"AMRESH",
		"amresh.terminal@gmail.com",
		24,
		"Male",
	}
}


func (u User) testUser(w http.ResponseWriter, r http.Request) error {
return nil
}
