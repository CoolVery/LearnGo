package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	dayNow := time.Now().Weekday()
	switch dayNow {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	}
	return "Неверный ввод"
}

func dayOrNight() string {
	hourNow := time.Now().Hour()
	if hourNow >= 10 && hourNow <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int {
	dayNow := time.Now().Weekday()
	result := time.Friday - dayNow
	return int(result)
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	dayNow := currentDayOfTheWeek()
	if strings.ToLower(answer) == strings.ToLower(dayNow) {
		return true
	}
	return false
}

func СheckNowDayOrNight(answer string) (bool, error) {
	lenString := utf8.RuneCountInString(answer)
	if lenString > 4 || lenString < 4 {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	hourNow := dayOrNight()
	stringParse := strings.ToLower(answer)
	if stringParse == strings.ToLower(hourNow) {
		return true, nil
	}
	return false, nil
}

func main() {
	fmt.Println(currentDayOfTheWeek())
	fmt.Println(dayOrNight())
	fmt.Println(nextFriday())
}
