package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (person Person) PrettyPrint() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", person.name, person.age, person.address)
}
