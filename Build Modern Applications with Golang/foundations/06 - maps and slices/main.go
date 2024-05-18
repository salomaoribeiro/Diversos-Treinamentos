package main

import (
	"log"
	"slices"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	myMapI := make(map[string]int)
	myMapMe := make(map[string]User)

	myMap["dog"] = "Samson"
	myMapI["dog"] = 5
	myMap["cat"] = "Cassie"

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMapMe["me"] = me

	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	myMap["dog"] = "New Name"
	log.Println(myMap["dog"])
	log.Println(myMapI["dog"])

	log.Println(myMapMe["me"])

	var mySlice []string

	mySlice = append(mySlice, "Salomão")
	mySlice = append(mySlice, "Kalleb")
	mySlice = append(mySlice, "Fletcher")
	mySlice = append(mySlice, "William")

	sliceChar := []string{"a", "j", "d", "k", "b"}

	log.Println(slices.Sort(mySlice))
	log.Println(slices.Sort(sliceChar))
}
