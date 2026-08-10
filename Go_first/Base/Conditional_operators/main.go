package main

import (
	"fmt"
)

func main() {
	// Задача 1
	// var language string
	// fmt.Scanln(&language)
	// if language == "Go" {
	// 	fmt.Println("Go!")
	// } else {
	// 	fmt.Println("Я знаю только Go!")
	// }

	// Задача 2
	// var firstNum, secondNum int
	// fmt.Scan(&firstNum, &secondNum)
	// if firstNum > secondNum {
	// 	fmt.Println("Первое число больше второго")
	// } else if firstNum < secondNum {
	// 	fmt.Println("Второе число больше первого")
	// } else {
	// 	fmt.Println("Числа равны")
	// }

	//Задача 3
	// var rub, kops, sum int
	// fmt.Scan(&rub, &kops, &sum)
	// rub += kops / 100
	// if rub >= sum {
	// 	fmt.Println("Сегодня будет вкусный кофе!")
	// } else {
	// 	fmt.Println("Стоит подкопить")
	// }

	//Задача 4
	// var num float64
	// fmt.Scanln(&num)
	// if sqrtValue := math.Sqrt(num); !math.IsNaN(sqrtValue) {
	// 	multiplier := math.Pow(10, 3)
	// 	fmt.Printf("%.3f",math.Round(sqrtValue*multiplier) / multiplier)
	// } else {
	// 	fmt.Println("-1")
	// }

	//Задача 5
	var num int
	fmt.Scanln(&num)
	if num == 0 {
		fmt.Println("Число 0")
	} else if value := num / 10; value == 0 {
		fmt.Println("Число однозначное")
	} else if num % 2 == 0 {
		fmt.Println("Число чётное")
	} else if num > 0 {
		fmt.Println("Число положительное")
	} else {
		fmt.Println("Число красивое")
	}
}