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

func getUser(w http.ResponseWriter, r http.Request) User {
	return User{
		1,
		"AMRESH",
		"amresh.terminal@gmail.com",
		24,
		"Male",
	}
}


func testUser(w http.ResponseWriter, r http.Request) User {
	return User{
		2,
		"AMRESH",
		"amresh.terminal@gmail.com",
		24,
		"Male",
	}
}
