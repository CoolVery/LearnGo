package main

import (
	"math"
	"slices"

	"github.com/CoolVery/LearnGo.git/classes"
	"github.com/CoolVery/LearnGo.git/interfaces"
)

func SortNums(nums []uint) {
	slices.Sort(nums)
}

func SortNames(names []string) {
	slices.Sort(names)
}

func SortAndMerge(left, right []int) []int {
	slices.Sort(left)
	slices.Sort(right)
	resultSlice := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			resultSlice = append(resultSlice, left[i])
			i++
		} else {
			resultSlice = append(resultSlice, right[j])
			j++
		}
	}

	// Дописываем оставшиеся элементы, если они есть.
	if i < len(left) {
		resultSlice = append(resultSlice, left[i:]...)
	}
	if j < len(right) {
		resultSlice = append(resultSlice, right[j:]...)
	}

	return nil

}

func main() {
	var company interfaces.CompanyInterface = &classes.Company{}
	company.AddWorkerInfo("ddd", "директор", 2, 2)
	company.AddWorkerInfo("ddd", "зам. директора", 1, 1)
	company.AddWorkerInfo("ddd", "директор", 20, 20)
	company.AddWorkerInfo("ddd", "директор", 5, 5)
	company.SortWorkers()
}