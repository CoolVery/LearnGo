package main

import (
	"fmt"
	"math"
	"unicode"
)

func myFunc() int {
	a := 1
	return a
}

func SquareRoots(a, b, c float64) (float64, float64) {
	discriminant := findDiscriminant(a, b, c)
	switch {
	case discriminant < 0:
		return 0, 0
	case discriminant == 0:
		xFirst := -b / (2 * a)
		return xFirst, xFirst
	case discriminant > 0:
		xFirst := (-b + math.Sqrt(discriminant)) / (2 * a)
		xSecond := (-b - math.Sqrt(discriminant)) / (2 * a)
		return xFirst, xSecond
	}
	return 0, 0
}

func findDiscriminant(a, b, c float64) float64 {
	var discriminant float64
	discriminant = math.Pow(b, 2) - 4*a*c
	return discriminant
}

func checkPassword(password string) bool {
	if hasMinimumLength(password, 8) && hasUpper(password) {
		return true
	}
	return false
}

func hasMinimumLength(password string, minLenght int) bool {
	if len(password) >= minLenght {
		return true
	}
	return false
}

func hasLowerCase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}
func hasUpper(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func ratePassword(password string) string {
	var countSuccess int
	if hasLowerCase(password) {
		countSuccess++
	}
	if hasUpper(password) {
		countSuccess++
	}
	if hasMinimumLength(password, 8) {
		countSuccess++
	}
	switch countSuccess {
	case 1:
		return "Слабый пароль"
	case 2:
		return "Средний пароль"
	case 3:
		return "Сложный пароль"
	}
	return "Пароль недостаточно безопасен, придумайте новый"
}

func main() {
	fmt.Println(hasMinimumLength("a", 10))
	fmt.Println(hasMinimumLength("", 0))
	fmt.Println(hasMinimumLength("aaa", 3))
	fmt.Println(hasMinimumLength("a", 1))
	fmt.Println(hasMinimumLength("afaghlonfwdfbomwdfbj", 10))

	fmt.Println(hasUpper("A"))
	fmt.Println(hasUpper(""))
	fmt.Println(hasUpper("f"))
	fmt.Println(hasUpper("wjefbhqjvoqnihbejbhefhqwlbhdw"))
	fmt.Println(hasUpper("AAAAAAAAAAAAAA"))
	fmt.Println(hasUpper("Aerfbjebrfh"))
	fmt.Println(hasUpper("gerfbjebrfhF"))
	fmt.Println(hasLowerCase("geaaarfbJebrfhl"))
}
