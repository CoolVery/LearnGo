package main

import (
	"reflect"
	"fmt"
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
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0
	
	travelTimes := EstimateTravelTime(vehicles, distance)

	fmt.Printf("Ожидается время для автомобиля %.2f часа\n", travelTimes["classes.Car"])
	fmt.Printf("Ожидается время для мотоцикла %.2f часа", travelTimes["classes.Motorcycle"])
}
