package main

import (
	"fmt"
	"strings"
	//"math"
)

func main() {
	//Задача 1
	// var value string
	// for i := 0; i < 5; i++ {
	// 	fmt.Scanln(&value)
	// 	if value == "Go" {
	// 		fmt.Println("Go!")
	// 	} else {
	// 		fmt.Println("Я знаю только Go!")
	// 	}
	// }

	//Задача 2
	// var countStudent int
	// var score float64
	// fmt.Scanln(&countStudent)
	// for i := 0; i < countStudent; i++ {
	// 	fmt.Scanln(&score)
	// 	switch {
	// 	case score >= 90 && score <= 100:
	// 		fmt.Println(5)
	// 	case score >= 75 && score <= 89:
	// 		fmt.Println(4)
	// 	case score >= 50 && score <= 74:
	// 		fmt.Println(3)
	// 	case score >= 0 && score <= 49:
	// 		fmt.Println(2)
	// 	default:
	// 		fmt.Println("Неверный балл")
	// 	}
	// }

	//Задача 3
	// var countProduct, discount int
	// var count, sum float64
	// fmt.Scanln(&countProduct)
	// fmt.Scanln(&discount)
	// for i := 0; i < countProduct; i++ {
	// 	fmt.Scanln(&count)
	// 	sum += count
	// }
	// discountValue := sum * float64(discount) / 100
	// allSum := sum - discountValue
	// if allSum - math.Trunc(allSum) == 0 {
	// 	fmt.Println(allSum)
	// } else {
	// 	fmt.Printf("%.3f", allSum)
	// }

	// //Задача 4
	// var value string
	// for {
	// 	fmt.Scanln(&value)
	// 	if value == "да" || value == "нет" || value == "чёрный" || value == "белый" {
	// 		fmt.Println("Поражение")
	// 		break
	// 	} else {
	// 		fmt.Println("Игра продолжается")
	// 	}
	// }

	//Задача 5
	var countWord, checkDoubleWord int
	var stringValue, checkWord string
	fmt.Scanln(&countWord)
	fmt.Scanln(&checkWord)
	fmt.Scanln(&stringValue)
	for i := 0; i < countWord; i++ {		
		fmt.Scan(&stringValue)
		if checkWord == strings.ToLower(stringValue) {
			checkDoubleWord++
		} 
	}
	fmt.Println(checkDoubleWord)
}
