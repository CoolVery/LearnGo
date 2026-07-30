package main

import "fmt"

type Dog struct {
}

func (dog Dog) MakeSound() {
	fmt.Println("Гав!")
}
