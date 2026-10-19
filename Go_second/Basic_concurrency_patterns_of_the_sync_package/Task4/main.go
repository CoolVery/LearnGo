package main

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)
type Student struct {
	Name string
	Score int
}
func CompareList(names []string) (map[string]string, error) {
	allStudents := make([]Student, 0, len(names))
	allNums := make([]int, 0, len(names))

	for _, n := range names {
		resp, err := http.Get("http://localhost:8082/mark?name=" + n)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusNotFound:
			return nil, errors.New("студент не найден")
		case http.StatusInternalServerError:
			return nil, errors.New("у сервера проблема")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
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
	mapAllStudents := make(map[string]string, 0)
	for _, s := range allStudents {
		if s.Score > average {
			mapAllStudents[s.Name] = ">"
		} else if s.Score < average {
			mapAllStudents[s.Name] = "<"
		} else {
			mapAllStudents[s.Name] = "="
		}
	}
	
	return mapAllStudents, nil
}

