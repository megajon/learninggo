package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	p := person{
		FirstName:  "Jonathan",
		MiddleName: stringp("Christopher"),
		LastName:   "Seubert",
	}
	fmt.Println(p.MiddleName)
}

func stringp(s string) *string {
	return &s
}
