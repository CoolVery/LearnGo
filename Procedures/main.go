package main

import "fmt"

func GoOrNot(value string) {
	if value == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}

func PrintFlightRow(numberRaise, numberRegister, numberGaite int, startCity, endCity string, isRegistery bool, continueFly float64) {
	switch isRegistery {
	case true:
		fmt.Printf("| %d | %s - %s | %d регистрация продолжается\n", numberRaise, startCity, endCity, numberRegister)
	case false:
		fmt.Printf("| %d | %s - %s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |\n", numberRaise, startCity, endCity, numberGaite, continueFly)
	}
}

func main() {
	// GoOrNot("Go")
	PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, false)
	PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, true)
}
