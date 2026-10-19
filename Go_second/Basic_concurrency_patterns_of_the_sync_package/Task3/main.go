package main

import (
	"errors"
	"io"
	"net/http"
	"slices"
	"strconv"
	"strings"
)
type Student struct {
	Name string
	Score int
}
func BestStudents(names []string) (string, error) {
	allStudents := make([]Student, 0, len(names))
	allNums := make([]int, 0, len(names))

	for _, n := range names {
		resp, err := http.Get("http://localhost:8082/mark?name=" + n)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusNotFound:
			return "", errors.New("студент не найден")
		case http.StatusInternalServerError:
			return "", errors.New("у сервера проблема")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		
		num, _ := strconv.Atoi(string(body))
		student := Student{
			Name: n,
			Score: num,
		}
		allStudents = append(allStudents, student)
		allNums = append(allNums, num)
	}
	var average, allSum int
	for _, n := range allNums {
		allSum += n
	}
	average = allSum / len(allNums)
	sliceStudentsWithUpScore := make([]Student, 0)
	for _, s := range allStudents {
		if s.Score > average {
			sliceStudentsWithUpScore = append(sliceStudentsWithUpScore, s)
		}
	}
	slices.SortFunc(sliceStudentsWithUpScore, func(a, b Student) int {
		return strings.Compare(a.Name, b.Name)
	})
	
	sliceStrName := make([]string, 0, len(sliceStudentsWithUpScore))
	for _, s := range sliceStudentsWithUpScore {
		sliceStrName = append(sliceStrName, s.Name)
	}
	result := strings.Join(sliceStrName, ",")
	return result, nil
}
