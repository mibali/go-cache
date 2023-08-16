package main

import (
	"fmt"

	"github.com/mcuadros/go-defaults"
)

type Person struct {
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	person := &Person{}
	defaults.SetDefaults(person)
	fmt.Printf("Default Person: %+v\n", person)

	// You can also assign values and use the package
	person.Name = "John"
	person.Age = 30
	person.Email = "john@example.com"
	person.Active = true

	fmt.Printf("Assigned Person: %+v\n", person)
}
