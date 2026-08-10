package classes

import (
	"fmt"
)

type animal struct {
	name (string)
	species (string)
	age (int)
	sound (string) 
}

func NewAnimal(name, species string, age int, sound string) animal {
	newAnimal := animal {
		name: name,
		species: species,
		age: age,
		sound: sound,
	}
	return newAnimal
}

func (animal animal) MakeSound() string {
	return animal.sound
}
func (animal animal) GetName() string {
	return animal.name
}
func (animal animal) GetInfo() string {
	resultString := fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", animal.name, animal.species, animal.age)
	return resultString
}

