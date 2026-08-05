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
	maxLen := math.Max(float64(len(left)), float64(len(right)))
	for i := 0; i < int(maxLen); i++ {
		switch int(maxLen) == len(left) {
		case true:

		}
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