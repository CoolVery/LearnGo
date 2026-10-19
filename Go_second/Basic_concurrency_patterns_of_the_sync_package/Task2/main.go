package main

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

func Average(names []string) (int, error) {
	allNums := make([]int, 0, len(names))
	for _, n := range names {
		resp, err := http.Get("http://localhost:8082/mark?name=" + n)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusNotFound:
			return 0, errors.New("студент не найден")
		case http.StatusInternalServerError:
			return 0, errors.New("у сервера проблема")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		num, _ := strconv.Atoi(string(body))
		allNums = append(allNums, num)
	}
	var average, allSum int
	for _, n := range allNums {
		allSum += n
	}
	average = allSum / len(allNums)
	return average, nil
}
