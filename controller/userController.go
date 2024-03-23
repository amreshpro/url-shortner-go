package controller

import (
	"net/http"
)

type User struct {
	name  string
	email string
	age   int16
	sex   string
}

func getUser(w http.ResponseWriter, r http.Request) User {
	return User{
		"AMRESH",
		"amresh.terminal@gmail.com",
		24,
		"Male",
	}
}
