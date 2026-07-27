package main

import (
	"errors"
	"fmt"
	"strings"
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

var Balance = 0.0

func topUpBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("amount is incorrect")
	}
	Balance += amount
	return nil

}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("amount is incorrect")
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {

	errorValue := topUpBalance(amount)
	if Balance < 0 {
		return 0, errors.New("balance is incorrect")
	}
	if errorValue != nil {
		Balance = 0
	}
	return Balance, errorValue
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	errorValue := chargeFromBalance(amount)
	if Balance < 0 {
		return 0, errors.New("balance is incorrect")
	}
	if errorValue != nil {
		Balance = 0
	}
	return Balance, errorValue
}

func CheckLetters(text string) string {
	countRuneE := strings.Count(text, "е")
	if countRuneE == 0 {
		return "Текст готов к публикации!"
	}
	result := fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст.", countRuneE)
	return result
}

func PrintComplexNumber(z complex64) {
	realPart := real(z)
	imagPart := imag(z)
	fmt.Printf("Действительная часть: %.2f. Мнимая часть: %.2f", realPart, imagPart)
}

func main() {

	fmt.Println(CheckLetters("dsёевыё"))
}
