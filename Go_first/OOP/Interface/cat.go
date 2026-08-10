package main

import "fmt"

type Cat struct {
}

func (cat Cat) MakeSound() {
	fmt.Println("Мяу!")
}
