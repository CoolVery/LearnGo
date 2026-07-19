package main

import (
    "fmt"
)

func main() {
    // Задача 1
	// var firstNum, secondNum, thirdNum float64
    // fmt.Scan(&firstNum, &secondNum, &thirdNum)
    // if firstNum == secondNum && secondNum == thirdNum && thirdNum == firstNum {
    //     fmt.Println("Максимальное равенство")
    // } else {
    //     fmt.Println("Не равны")
    // }
	
	// Задача 2
	// var firstPassword, secondPassword string
	// fmt.Scanln(&firstPassword)
	// fmt.Scanln(&secondPassword)
	// lenFirstPassword := len(firstPassword)
	// lenSecondPassword := len(secondPassword)
	// if lenFirstPassword >= 8 && lenSecondPassword >= 8 {
	// 	fmt.Println("Оба пароли надёжные")
	// } else if lenFirstPassword >= 8 {
	// 	fmt.Println("Только первый пароль надёжны")
	// } else if lenSecondPassword >= 8 {
	// 	fmt.Println("Только второй пароль надёжный")
	// } else {
	// 	fmt.Println("Оба пароля ненадёжные")
	// }
    //Задача 3
	// var firstPlayer, secondPlayer string
	// fmt.Scanln(&firstPlayer)
	// fmt.Scanln(&secondPlayer)
    // if firstPlayer == secondPlayer {
	// 	fmt.Println("Ничья")
	// 	return
	// }
	// switch firstPlayer {
	// case "камень":
	// 	if secondPlayer == "бумага" {
	// 		fmt.Println("Второй игрок победил")
	// 	} else {
	// 		fmt.Println("Первый игрок победил")
	// 	}
	// case "бумага":
	// 	if secondPlayer == "ножницы" {
	// 		fmt.Println("Второй игрок победил")
	// 	} else {
	// 		fmt.Println("Первый игрок победил")
	// 	}
	// case "ножницы":
	// 	if secondPlayer == "камень" {
	// 		fmt.Println("Второй игрок победил")
	// 	} else {
	// 		fmt.Println("Первый игрок победил")
	// 	}
	// }

	var sign string
	var valueNum int
	fmt.Scan(&sign, &valueNum)
	
	if valueNum == 0 {
		fmt.Println("Стоит надеть куртку")
        return
	}
	switch sign {
	case "+":
		if valueNum > 20 {
			fmt.Println("Стоит надеть майку и шорты")
		} else if valueNum >= 10 {
			fmt.Println("Стоит надеть штаны и кофту")
		} else {
			fmt.Println("Стоит надеть куртку")
		}
	case "-":
		if valueNum <= 5 {
			fmt.Println("Стоит надеть куртку")
		} else {
			fmt.Println("Стоит надеть зимнюю куртку")
		}
	}
}