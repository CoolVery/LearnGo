package main

import (
	"reflect"
	"fmt"
	"strings"
	. "polym/interfaces"
	. "polym/classes"
	
)

func  EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	resultMap := make(map[string]float64)
	for _, value := range vehicles {
		typeName := reflect.TypeOf(value).String()
		time := value.CalculateTravelTime(distance)
		resultMap[typeName] = time
	}
	
	return resultMap
}

func main() {
	// car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	// motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	// vehicles := []Vehicle{car, motorcycle}

	// distance := 200.0
	
	// travelTimes := EstimateTravelTime(vehicles, distance)

	// fmt.Printf("Ожидается время для автомобиля %.2f часа\n", travelTimes["classes.Car"])
	// fmt.Printf("Ожидается время для мотоцикла %.2f часа", travelTimes["classes.Motorcycle"])

	urw := &UpperReaderWriter{}

	testData := []byte("I love Golang")
	bytesWritten := urw.Write(testData)
	fmt.Printf("Записано %d байт: %s\n", bytesWritten, testData)

	testData = []byte("Обожаю Яндекс")
	bytesWritten = urw.Write(testData)
	fmt.Printf("Записано %d байт: %s\n", bytesWritten, testData)

	readData := urw.Read()
	fmt.Printf("Прочитано: %s\n", readData)

	if urw.UpperString != strings.ToUpper(string(testData)) {
		fmt.Println("Ошибка: строка не преобразована в верхний регистр")
	}

	var _ Reader = urw
	var _ Writer = urw
	var _ ReaderWriter = urw
}
