package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	BirthDate   time.Time
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         string
}

func main() {
	s2 := "six"
	log.Println("s is", s)
	log.Println("s2 is", s2)

	log.Println(saySomething("xxx"))

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName, user.LastName)
	log.Println("BirthDate:", user.BirthDate)
}

func saySomething(s string) (string, string) {
	return s, "World"
}
