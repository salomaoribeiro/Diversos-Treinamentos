package main

import "fmt"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myvar myStruct
	myvar.FirstName = "Name 1"

	myvar2 := myStruct{
		FirstName: "Name 2",
	}

	fmt.Println(myvar.FirstName)
	fmt.Println(myvar2.FirstName)

	fmt.Println("printing with func", myvar.printFirstName())
	fmt.Println("printing with func", myvar2.printFirstName())
}
