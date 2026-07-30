package main

import (
	"fmt"
	"slices"
	"sort"
	"strings"
)

func SplitStringsToWords(text string) []string {
	tempString := ""
	lenString := len(text)
	splitStringSlice := make([]string, 0)
	targetRune := []rune{' ', '.', ',', '!', '?'}
	for index, runeValue := range text {
		if slices.Contains(targetRune, runeValue) {
			if index+1 != lenString && slices.Contains(targetRune, rune(text[index+1])) {
				continue
			}
			splitStringSlice = append(splitStringSlice, tempString)
			tempString = ""
		} else {
			tempString += string(runeValue)
		}

	}
	if len(tempString) != 0 {
		splitStringSlice = append(splitStringSlice, tempString)
	}
	return splitStringSlice
}

func CountUniqueWords(sliceStrings []string) map[string]int {
	resultMap := make(map[string]int)
	for _, value := range sliceStrings {
		value = strings.ToLower(value)
		resultMap[value]++
	}
	return resultMap
}

func FoundMaxWord(mapWords map[string]int) (string, int) {
	maxValue := 0
	resultString := ""
	for value, _ := range mapWords {
		if mapWords[value] > maxValue {
			maxValue = mapWords[value]
			resultString = value
		}
	}
	return resultString, maxValue
}

func getTopWords(wordMap map[string]int, n int) []string {
	resultSlice := make([]string, 0)
	foundWord := 0
	tempSlice := make([]int, 0)
	tempMap := make(map[int]string)
	for value, _ := range wordMap {
		tempSlice = append(tempSlice, wordMap[value])
		tempMap[wordMap[value]] = value
	}
	sort.Slice(tempSlice, func(i, j int) bool {
		return tempSlice[i] > tempSlice[j]
	})
	for {
		for key, value := range tempMap {
			if foundWord == n {
				return resultSlice
			}
			if key == tempSlice[foundWord] {
				resultSlice = append(resultSlice, value)
				foundWord++
			}
		}
	}
}

func AnalyzeText(text string) {
	sliceToWork := SplitStringsToWords(text)
	fmt.Printf("Количество слов: %d\n", len(sliceToWork))
	mapToSliceToWork := CountUniqueWords(sliceToWork)
	maxUniaqyeWord, maxUniqueCount := FoundMaxWord(mapToSliceToWork)
	fmt.Printf("Количество уникальных слов: %d\n", len(mapToSliceToWork))
	fmt.Printf("Самое часто встречающееся слово: %s (встречается %d раз)\n", maxUniaqyeWord, maxUniqueCount)
	fmt.Printf("Топ-5 самых часто встречающихся слов:\n")
	topWordToSlice := getTopWords(mapToSliceToWork, 5)
	for _, value := range topWordToSlice {
		fmt.Printf("\"%s\" : %d раз\n", value, mapToSliceToWork[value])
	}
}

func main() {
	AnalyzeText("one two two three three three four four four four five five five five five")
	AnalyzeText("Go очень очень очень ОЧЕНЬ ОчЕнь очень оЧЕНь классный классный! go просто, ну просто классный. GO Классный!")
	AnalyzeText("Я так люблю море. Я на море. Я так люблю. Море! Я море!!! ЛЮБЛЮ МОРЕ")
}
