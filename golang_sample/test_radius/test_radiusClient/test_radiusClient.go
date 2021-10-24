package main

import (
	"log"

	"github.com/kirves/goradius"
)

// var (
// 	server string = "192.168.3.2"
// 	port   string = "1812"
// 	secret string = "testing123"
// 	// timeout  time.Duration = 5
// 	username string = "bob"
// 	password string = "hello"
// 	nasId    string = ""
// )
// var (
// 	server string = "192.168.3.2"
// 	port   string = "1812"
// 	secret string = "testing123"
// 	// timeout  time.Duration = 5
// 	username string = "alice"
// 	password string = "passme"
// 	nasId    string = ""
// )
//Test for window server
var (
	server string = "192.168.3.2"
	port   string = "1812"
	secret string = "testing123"
	// timeout  time.Duration = 5
	username string = "bob"
	password string = "hello"
	nasId    string = ""
)

func TestGoradius() {
	auth := goradius.Authenticator(server, port, secret)
	ok, err := auth.Authenticate(username, password, nasId)
	if err != nil {
		panic(err)
		return
	}
	if ok {
		// user successfully authenticated
		log.Println("login in successfully")
		return
	}
}

func main() {
	TestGoradius()
}
