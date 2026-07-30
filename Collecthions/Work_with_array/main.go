package main

import "fmt"

func FiveSteps(array [5]int) [5]int {
	var resultArray [5]int
	var index = 0
	for i := 4; i >= 0; i-- {
		resultArray[index] = array[i]
		index++
	}
	return resultArray
}

func ThirdElementInArray(array [6]int) int {
	return array[2]
}

func FindMaxMinInArray(array [10]int) (int, int) {
	var resultMin, resultMax = array[0], array[1]
	for _, value := range array {
		if value > resultMax {
			resultMax = value
		}
		if value < resultMin {
			resultMin = value
		}
	}
	return resultMax, resultMin
}

func SumOfArray(array [6]int) int {
	var resultSum int
	for _, value := range array {
		resultSum += value
	}
	return resultSum
}

func PrettyArrayOutput(array [9]string) {
	for index, value := range array {
		if index < 7 {
			fmt.Printf("%d я уже сделал: %s", index+1, value)
		} else {
			fmt.Printf("%d я уже сделал: %s", index+1, value)
		}
	}
}

func main() {

}
