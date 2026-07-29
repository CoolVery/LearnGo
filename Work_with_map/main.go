package main

import (
	"unicode/utf8"
)

func countOccurrences(nums []int) map[int]int {
	occurrences := make(map[int]int)

	for _, num := range nums {
		occurrences[num]++
	}

	return occurrences
}

func FindMaxKey(m map[int]int) int {
	maxValue := 0
	resultValue := 0
	for value, _ := range m {
		if m[value] > maxValue {
			resultValue = value
			maxValue = m[value]
		}
	}
	return resultValue
}

func SumOfValuesInMap(m map[int]int) int {
	resultSum := 0
	for value, _ := range m {
		resultSum += m[value]
	}
	return resultSum
}

func SwapKeysAndValues(m map[string]string) map[string]string {
	resultMap := make(map[string]string)
	for value, _ := range m {
		tempValue := m[value]
		resultMap[tempValue] = value
	}
	return resultMap
}

func CountingSort(contacts []string) map[string]int {
	resultMap := make(map[string]int)
	for _, value := range contacts {
		resultMap[value]++
	}
	return resultMap
}

func DeleteLongKeys(m map[string]int) map[string]int {
	resultMap := make(map[string]int)
	for value, _ := range m {
		if utf8.RuneCountInString(value) <= 6 {
			resultMap[value] = m[value]
		}
	}
	return resultMap
}

func main() {
	m := map[string]string{
		"Яндекс":        "+74957397000",
		"Музей Яндекса": "+74991101133",
	}
	SwapKeysAndValues(m)
}
