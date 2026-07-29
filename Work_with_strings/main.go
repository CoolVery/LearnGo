package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func CountLengthAndBytes(first, second string) string {
	allString := first + second
	countByte := len(allString)
	countRune := utf8.RuneCountInString(allString)
	resultString := fmt.Sprintf("Объединённая строка: %s. Количество байт: %d. Количество символов: %d.", allString, countByte, countRune)
	return resultString
}

func CheckOnlyASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func NumbersToLetters(input string) string {
	var resiltString string
	firstLetter := []rune(input)
	for _, r := range firstLetter {
		currentRune := string(r)
		switch currentRune {
		case "0":
			resiltString += "ноль"
		case "1":
			resiltString += "один"
		case "2":
			resiltString += "два"
		case "3":
			resiltString += "три"
		case "4":
			resiltString += "четыре"
		case "5":
			resiltString += "пять"
		case "6":
			resiltString += "шесть"
		case "7":
			resiltString += "семь"
		case "8":
			resiltString += "восемь"
		case "9":
			resiltString += "девять"
		case "+":
			resiltString += "плюс"
		case "*":
			resiltString += "умножить на"
		case "-":
			resiltString += "минус"
		case "/":
			resiltString += "разделить на"
		default:
			resiltString += currentRune
		}
	}
	return resiltString
}

func main() {
	fmt.Println(NumbersToLetters("(1 + 2) * 3 / 8"))
}
