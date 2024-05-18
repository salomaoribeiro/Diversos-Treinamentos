package main

import "log"

func main() {
	myString := "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)

	log.Println("new value of myString", myString)

	changeWithoutPointer(myString)

	log.Println("value of myString without pointer", myString)
}

func changeUsingPointer(s *string) {
	newValue := "Red"

	*s = newValue
}

func changeWithoutPointer(s string) {
	newValue := "Yellow"
	s = newValue
}
