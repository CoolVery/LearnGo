package main

import (
	"fmt"
	"strconv"
)

func main() {
	var stringValue, intStringValue string
	var (
		firstStudent  = "-"
		secondStudent = "-"
		thiardStudent = "-"
		fourStudent   = "-"
		fifthStudent  = "-"
	)
	var countInStack int
	for {
		stringValue, intStringValue = "", ""
		fmt.Scanln(&stringValue, &intStringValue)
		// stringValue = "очередь"
		if len(stringValue) != 0 || len(intStringValue) != 0 {
			intValue, err := strconv.Atoi(intStringValue)
			if err != nil {

			}
			if (intValue < 1 || intValue > 5) && err == nil {
				fmt.Printf("Запись на место номер %d невозможна: некорректный ввод\n", intValue)
				continue
			}
			switch stringValue {
			case "конец":
				fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", firstStudent, secondStudent, thiardStudent, fourStudent, fifthStudent)
				return
			case "очередь":
				fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", firstStudent, secondStudent, thiardStudent, fourStudent, fifthStudent)
			case "количество":
				fmt.Printf("Осталось свободных мест: %d\n", 5-countInStack)
				fmt.Printf("Всего человек в очереди: %d\n", countInStack)
			default:
				if firstStudent != "-" && secondStudent != "-" && thiardStudent != "-" && fourStudent != "-" && fifthStudent != "-" {
					fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", intValue)
				} else {
					switch intValue {
					case 1:
						if firstStudent == "-" {
							firstStudent = stringValue
							countInStack++
						} else {
							fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", intValue)
						}
					case 2:
						if secondStudent == "-" {
							secondStudent = stringValue
							countInStack++
						} else {
							fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", intValue)
						}
					case 3:
						if thiardStudent == "-" {
							thiardStudent = stringValue
							countInStack++
						} else {
							fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", intValue)
						}
					case 4:
						if fourStudent == "-" {
							fourStudent = stringValue
							countInStack++
						} else {
							fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", intValue)
						}
					case 5:
						if fifthStudent == "-" {
							fifthStudent = stringValue
							countInStack++
						} else {
							fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", intValue)
						}
					}
				}
			}
		}
	}
}
