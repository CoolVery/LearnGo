package main

import (
	"errors"
	"io"
	"net/http"
	"strconv"
)

func Compare(name1, name2 string) (string, error) {
	respFirst, err := http.Get("http://localhost:8082/mark?name=" + name1)
	defer respFirst.Body.Close()
	if err != nil {
		return "", err
	}
	if respFirst.StatusCode == 404 {
		return "", errors.New("студент не найден")
	}
	if respFirst.StatusCode == 500 {
		return "", errors.New("у сервера проблема")
	}
	respSecond, err := http.Get("http://localhost:8082/mark?name=" + name2)
	defer respSecond.Body.Close()
	if err != nil {
		return "", err
	}
	if respSecond.StatusCode == 404 {
		return "", errors.New("студент не найден")
	}
	if respSecond.StatusCode == 500 {
		return "", errors.New("у сервера проблема")
	}
	var numFirst, numSecond int
	
	body, err := io.ReadAll(respFirst.Body)
	if err != nil {
		return "", err
	}
	numFirst, _ = strconv.Atoi(string(body))

	body, err = io.ReadAll(respSecond.Body)
	if err != nil {
		return "", err
	}
	numSecond, _ = strconv.Atoi(string(body))
	
	if numFirst > numSecond {
		return ">", nil
	} else if numFirst < numSecond {
		return "<", nil
	} else {
		return "=", nil
	}
}