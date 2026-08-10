package main

import "fmt"

func GoOrNot(value string) {
	if value == "Go" {
		fmt.Println("Go!")
	} else {
		fmt.Println("Я знаю только Go!")
	}
}

func PrintFlightRow(numberRaise, startCity, endCity string, continueFly float64, numberRegister, numberGaite int, isRegistery bool) {
	switch isRegistery {
	case true:
		fmt.Printf("| %s | %s - %s | %d регистрация продолжается\n", numberRaise, startCity, endCity, numberRegister)
	case false:
		fmt.Printf("| %s | %s - %s | регистрация закончилась, проходите к гейту: %d | длительность полёта %.1f часа |\n", numberRaise, startCity, endCity, numberGaite, continueFly)
	}
}

func printPrice(count int, nameProduct string) {
	fmt.Printf("%s будет стоить %d рублей\n", nameProduct, count)
}

func BuyFries(size string) {
	switch size {
	case "S":
		printPrice(49, "Картошка фри")
	case "M":
		printPrice(89, "Картошка фри")
	case "L":
		printPrice(99, "Картошка фри")
	default:
		fmt.Println("Некорректный размер")
	}
}

func BuyCola(size string) {
	switch size {
	case "S":
		printPrice(99, "Кола")
	case "M":
		printPrice(119, "Кола")
	case "L":
		printPrice(139, "Кола")
	default:
		fmt.Println("Некорректный размер")
	}
}

func main() {
	// GoOrNot("Go")
	// PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, false)
	// PrintFlightRow("117B", "Москва", "Казань", 2, 115, 3, true)
	BuyCola("S")
}
