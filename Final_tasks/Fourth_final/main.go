package main

import (
	"fmt"
	 _ "strings"
	. "github.com/CoolVery/LearnGo.git/classes"
	. "github.com/CoolVery/LearnGo.git/interfaces"
)

func ZooShow(animals []Animal) {
	for _, value := range animals {
		infoAnimal := value.GetInfo()
		soundAnimal := value.MakeSound()
		fmt.Println(infoAnimal)
		fmt.Println(soundAnimal)
	}
}

func main() {
	// tiger := NewAnimal("Барсик", "Тигр", 5, "Ррр")

	// // Проверка GetInfo
	// expectedInfo := "Имя: Барсик, Вид: Тигр, Возраст: 5"
	// if info := tiger.GetInfo(); !strings.Contains(info, "Барсик") || !strings.Contains(info, "Тигр") {
	// 	fmt.Printf("GetInfo() = %v, ожидается %v\n", info, expectedInfo)
	// }

	// // Проверка MakeSound
	// if sound := tiger.MakeSound(); sound != "Ррр" {
	// 	fmt.Printf("MakeSound() = %v, ожидается 'Ррр'\n", sound)
	// }

	// // Проверка Name()
	// if name := tiger.GetName(); name != "Барсик" {
	// 	fmt.Printf("GetName() = %v, ожидается 'Барсик'\n", name)
	// }

	tiger := NewAnimal("Барсик", "Тигр", 5, "Ррр")
	penguin := NewAnimal("Пиня", "Пингвин", 2, "Кря")
	ZooShow([]Animal{tiger, penguin})

	keeper := ZooKeeper{}
	keeper.Feed(tiger)
}