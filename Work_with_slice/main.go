package main

import (
	"errors"
)

func SliceCopy(nums []int) []int {
	lenResult := len(nums)
	resultSlice := make([]int, lenResult)
	for i := 0; i < lenResult; i++ {
		resultSlice[i] = nums[i]
	}
	return resultSlice
}

func Join(nums1, nums2 []int) []int {
	lenFirst, lenSecond := len(nums1), len(nums2)
	capFirst, capSecond := cap(nums1), cap(nums2)
	resultSlice := make([]int, capFirst+capSecond)
	for i := 0; i < lenFirst; i++ {
		resultSlice[i] = nums1[i]
	}
	for i := 0; i < lenSecond; i++ {
		resultSlice[lenFirst+i] = nums2[i]
	}
	return resultSlice
}

func Mix(nums []int) []int {
	lenResult := len(nums)
	indexResult := 0
	partLenResult := lenResult / 2
	resultSlice := make([]int, lenResult)
	for i := 0; i < partLenResult; i++ {
		resultSlice[indexResult] = nums[i]
		indexResult++
		resultSlice[indexResult] = nums[i+partLenResult]
		indexResult++
	}
	return resultSlice
}

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("Стоимость не может быть меньше 0")
	}
	if nums == nil {
		return nil, errors.New("Список не может быть пустым")
	}
	resultSlice := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < limit {
			resultSlice = append(resultSlice, nums[i])
		}
		if len(resultSlice) == 5 {
			break
		}
	}
	return resultSlice, nil
}

func main() {
	Mix([]int{0, 1, 2, 3, 4, 5})
}
