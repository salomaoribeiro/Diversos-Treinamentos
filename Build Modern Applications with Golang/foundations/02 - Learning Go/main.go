package main

import "fmt"

func main() {
	fmt.Println("Hello World.")

	var whatsToSay string
	var i int

	whatsToSay = "Goodbye, cruel world"
	fmt.Println(whatsToSay)

	i = 7

	fmt.Println("i is set to", i)
	fmt.Println(saySomething())

	something, otherThing := saySomething()
	fmt.Println(something, otherThing)
}

func saySomething() (string, string) {
	return "something", "else"
}
