package main

import (
	"errors"
	"fmt"
)

var (
	ErrScanln       = errors.New("Ошибка: некорректный ввод")
	ErrAgeInput     = errors.New("Ошибка: невалидный возраст")
	ErrLenghtName   = errors.New("Ошибка: невалидное имя")
	ErrPasportInput = errors.New("Ошибка: невалидная серия и номер паспорта")
)

var (
	DivisionByZeroError = errors.New("division by zero is not allowed")
)

func firstTask() {
	var age int
	var name string
	var passportSeriesAndNumber string

	countValues, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil {
		fmt.Println(ErrScanln)
		return
	}
	if countValues == 3 {
		if age < 14 || age > 150 {
			fmt.Println(ErrAgeInput)
			return
		}
		if len(name) < 2 {
			fmt.Println(ErrLenghtName)
			return
		}
		if len(passportSeriesAndNumber) != 10 {
			fmt.Println(ErrPasportInput)
			return
		}
		fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
	}
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError
	}
	result := float64(a) / float64(b)
	return result, nil
}

func main() {

	fmt.Println(Divide(15, 2))
}
