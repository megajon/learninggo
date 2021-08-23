package main

import "fmt"

func main() {
	people := []Person{
		{"Pat", "Paterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
